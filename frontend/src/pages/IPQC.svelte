<script>
</script>

<main class="h-full bg-gray-50 p-1.5">
  <section class="bg-white rounded-lg shadow-sm p-2 mb-2">
    <div
      class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-2"
    >
      <!-- 基本信息 -->
      <div class="flex items-center space-x-2">
        <label
          for="ipqcNumber"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >检验单号:</label
        >
        <input
          id="ipqcNumber"
          type="text"
          placeholder="IPQC单据号"
          required
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
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
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      <div class="flex items-center space-x-2">
        <label
          for="lotNumber"
          class="w-20 text-right text-xs font-medium text-gray-700 whitespace-nowrap"
          >批次数量:</label
        >
        <select
          id="lotNumber"
          class="flex-1 px-2 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="成型">成型</option>
          <option value="电镀">电镀</option>
          <option value="阻焊">阻焊</option>
          <option value="图形">图形</option>
          <option value="表面处理">表面处理</option>
        </select>
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
</main>
