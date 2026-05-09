<script lang="ts">
  import { SearchIpqcByCheckCode } from "../../wailsjs/go/main/Ipqc.js";
  import type { models } from "../../wailsjs/go/models.js";

  // 使用 Svelte 5 的 runes 语法声明响应式状态
  let ipqcNumber = $state("");
  let lotId = $state("");
  let productModel = $state("");
  let lotNumber = $state("");
  let checkNumber = $state(0);
  let result = $state("");
  let ncnNumber = $state("");
  let remark = $state("");

  let searchResults = $state<models.IPqcCheckTask[]>([]);
  let showDropdown = $state(false);
  let isLoading = $state(false);

  // 搜索IPQC单据号
  async function searchIpqcNumbers(query: string) {
    if (!query || query.length < 6) {
      searchResults = [];
      showDropdown = false;
      return;
    }

    isLoading = true;
    try {
      // 调用后端API搜索IPQC单据号
      const response = await SearchIpqcByCheckCode(query);
      searchResults = response || [];
      showDropdown = searchResults.length > 0;
    } catch (error) {
      console.error("搜索IPQC单据号失败:", error);
      searchResults = [];
      showDropdown = false;
    } finally {
      isLoading = false;
    }
  }

  // 处理输入变化
  function handleInputChange(event: Event) {
    const target = event.target as HTMLInputElement;
    ipqcNumber = target.value;
    searchIpqcNumbers(ipqcNumber);
  }

  // 选择搜索结果 - 使用 Svelte 5 语法
  function selectIpqcNumber(selectedData: models.IPqcCheckTask) {
    ipqcNumber = selectedData.CheckCode;
    lotId = selectedData.LotId;
    productModel = selectedData.ProductNumberVer;
    lotNumber = selectedData.LotNumber;
    checkNumber = selectedData.CheckNumber;
    ncnNumber = selectedData.NcnCode;
    remark = selectedData.Remark;
    showDropdown = false;
    searchResults = [];
  }

  // 处理输入框获得焦点
  function handleFocus() {
    if (ipqcNumber && ipqcNumber.length >= 6) {
      searchIpqcNumbers(ipqcNumber);
    }
  }

  // 点击外部关闭下拉列表
  function handleClickOutside(event: MouseEvent) {
    const target = event.target as Element;
    if (!target.closest("#ipqcNumber") && !target.closest(".dropdown")) {
      showDropdown = false;
    }
  }
</script>

<svelte:window on:click={handleClickOutside} />

<main class="h-full bg-gray-50 p-1.5">
  <section class="bg-white rounded-lg shadow-sm p-2 mb-2">
    <div
      class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-2"
    >
      <!-- 基本信息 -->
      <div class="flex items-center space-x-2 relative">
        <label
          for="ipqcNumber"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >检验单号:</label
        >
        <div class="flex-1 relative">
          <input
            id="ipqcNumber"
            type="text"
            bind:value={ipqcNumber}
            oninput={handleInputChange}
            onfocus={handleFocus}
            placeholder="输入IPQC单号(至少6位)"
            required
            class="w-full px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />

          <!-- 加载状态指示器 -->
          {#if isLoading}
            <div class="absolute inset-y-0 right-0 flex items-center pr-2">
              <div
                class="animate-spin rounded-full h-4 w-4 border-b-2 border-blue-600"
              ></div>
            </div>
          {/if}

          <!-- 搜索下拉列表 -->
          {#if showDropdown && searchResults.length > 0}
            <div
              class="dropdown absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-y-auto"
            >
              {#each searchResults as result}
                <div
                  role="button"
                  tabindex="0"
                  class="px-3 py-2 hover:bg-gray-100 cursor-pointer text-sm border-b border-gray-100 last:border-b-0"
                  onclick={() => selectIpqcNumber(result)}
                  onkeydown={(e) => {
                    if (e.key === "Enter" || e.key === " ") {
                      e.preventDefault();
                      selectIpqcNumber(result);
                    }
                  }}
                >
                  <div class="font-medium text-gray-900">
                    {result.CheckCode}
                  </div>
                  <div class="text-xs text-gray-500">
                    批次号: {result.LotId} | 产品型号: {result.ProductNumberVer}
                  </div>
                </div>
              {/each}
            </div>
          {:else if showDropdown && !isLoading && ipqcNumber.length >= 6}
            <div
              class="dropdown absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg"
            >
              <div class="px-3 py-2 text-sm text-gray-500">
                未找到匹配的IPQC单据号
              </div>
            </div>
          {/if}
        </div>
      </div>

      <div class="flex items-center space-x-2">
        <label
          for="lotId"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >批次号:</label
        >
        <input
          id="lotId"
          type="text"
          placeholder="批次号"
          bind:value={lotId}
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <div class="flex items-center space-x-2">
        <label
          for="productModel"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >产品型号:</label
        >
        <input
          id="productModel"
          type="text"
          placeholder="产品型号"
          bind:value={productModel}
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <div class="flex items-center space-x-2">
        <label
          for="lotNumber"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >批次数量:</label
        >
        <input
          id="lotNumber"
          type="text"
          placeholder="批次数量"
          bind:value={lotNumber}
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <div class="flex items-center space-x-2">
        <label
          for="checkNumber"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >检验数量:</label
        >
        <input
          id="checkNumber"
          type="text"
          placeholder="检验数量"
          bind:value={checkNumber}
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <div class="flex items-center space-x-2">
        <label
          for="result"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >判定结果:</label
        >
        <input
          id="result"
          type="text"
          placeholder="判定结果"
          bind:value={result}
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <div class="flex items-center space-x-2">
        <label
          for="NCNNumber"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >NCN单号:</label
        >
        <input
          id="NCNNumber"
          type="text"
          placeholder="NCN单号"
          bind:value={ncnNumber}
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <div class="flex items-center space-x-2">
        <label
          for="remark"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >备注:</label
        >
        <input
          id="remark"
          type="text"
          placeholder="备注"
          bind:value={remark}
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <!-- 搜索和重置按钮 -->
      <div class="flex flex-nowrap gap-4 justify-center">
        <button
          class="px-4 py-1 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors shadow-sm text-sm"
        >
          更新
        </button>
        <button
          class="px-4 py-1 bg-gray-600 text-white rounded-md hover:bg-gray-700 transition-colors shadow-sm text-sm"
        >
          重置
        </button>
      </div>
    </div>
  </section>
  <section class="bg-white rounded-lg shadow-sm p-2 mb-2">
    <!-- 这里可以放置IPQC相关的其他内容，例如检验记录列表等 -->
  </section>
</main>
