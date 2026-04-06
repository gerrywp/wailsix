<script>
  import { Tabs } from "@skeletonlabs/skeleton-svelte";
  import { X } from "@lucide/svelte"; // 关闭图标
  import { onMount } from "svelte";
  import Home from "./Home.svelte"; // 导入默认 tab 组件
  import NcnPage from "./NcnPage.svelte";
  import { EventsOn } from "../../wailsjs/runtime/runtime.js";

  let activeTab = $state("");
  let tabs = $state([]);

  // 初始化默认 tab
  onMount(() => {
    if (tabs.length === 0) {
      // 默认打开主页 tab
      const defaultTab = {
        id: "home",
        title: "主页",
        component: Home,
      };
      tabs = [defaultTab];
      activeTab = defaultTab.id;
    }
  });
  const routers = [
    { id: "home", title: "主页", component: Home },
    { id: "ncnpage", title: "NCN页面", component: NcnPage },
  ];
  EventsOn("navigate", (page) => {
    let tabToOpen = routers.find((r) => r.id === page);
    if (!tabToOpen) {
      return;
    }
    openTab(tabToOpen);
    //window.location.hash = page;
  });
  function openTab({ id, title, component }) {
    const exist = tabs.find((x) => x.id === id);
    if (exist) {
      activeTab = id;
      return;
    }
    tabs = [...tabs, { id, title, component }];
    activeTab = id;
  }

  function closeTab(e, id) {
    e.stopPropagation(); // 阻止触发 tab 切换
    if (tabs.length <= 1) {
      return;
    }

    const currentIndex = tabs.findIndex((t) => t.id === id);
    let newActiveTab = activeTab;

    if (activeTab === id) {
      // 如果关闭的是当前激活的 tab，需要确定新的激活 tab
      if (currentIndex > 0) {
        // 如果不是第一个 tab，激活前一个
        newActiveTab = tabs[currentIndex - 1].id;
      } else {
        // 如果是第一个 tab，激活下一个
        newActiveTab = tabs[currentIndex + 1].id;
      }
    }
    // 更新激活的 tab
    activeTab = newActiveTab;
    // 过滤掉要关闭的 tab
    tabs = tabs.filter((t) => t.id !== id);
  }
</script>

<!--Skeleton Tabs-->
<Tabs value={activeTab} onValueChange={(e) => (activeTab = e.value)}>
  <Tabs.List class="bg-surface-100 border-b border-surface-300 m-0">
    {#each tabs as tab (tab.id)}
      {#if tabs.some((t) => t.id === tab.id)}
        <Tabs.Trigger value={tab.id} class="flex items-center gap-2 px-1">
          {tab.title}
          <!-- 关闭按钮 -->
          <button
            onclick={(e) => closeTab(e, tab.id)}
            class="hover:bg-surface-300 rounded-full p-0.5 text-sm transition-colors"
          >
            <X size={14} />
          </button>
        </Tabs.Trigger>
      {/if}
    {/each}
    <Tabs.Indicator />
  </Tabs.List>

  <!-- 内容面板 -->
  {#each tabs as tab (tab.id)}
    <Tabs.Content value={tab.id}>
      <tab.component />
    </Tabs.Content>
  {/each}
</Tabs>
