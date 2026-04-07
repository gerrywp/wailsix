<script>
  import { X } from "@lucide/svelte";
  import { onMount, tick } from "svelte";
  import Home from "./Home.svelte";
  import NcnPage from "./NcnPage.svelte";
  import { EventsOn } from "../../wailsjs/runtime/runtime.js";

  let activeTab = $state("");
  let tabs = $state([]);

  const routers = [
    { id: "home", title: "主页", component: Home },
    { id: "ncnpage", title: "NCN页面", component: NcnPage },
  ];

  onMount(() => {
    if (tabs.length === 0) {
      const defaultTab = { id: "home", title: "主页", component: Home };
      tabs = [defaultTab];
      activeTab = defaultTab.id;
    }

    const off = EventsOn("navigate", (page) => {
      const tabToOpen = routers.find((r) => r.id === page);
      if (!tabToOpen) return;
      openTab(tabToOpen);
    });
    return off;
  });

  async function openTab({ id, title, component }) {
    const exist = tabs.find((x) => x.id === id);
    if (exist) {
      activeTab = id;
      return;
    }

    tabs = [...tabs, { id, title, component }];
    await tick();
    activeTab = id;
  }

  function closeTab(e, id) {
    e.stopPropagation();
    if (tabs.length <= 1) return;
    const idx = tabs.findIndex(t => t.id === id);
    if (idx === 0) return;

    // 切换激活项
    if (activeTab === id) {
      activeTab = tabs[idx - 1].id;
    }

    tabs = tabs.filter(t => t.id !== id);
  }
</script>
<div class="w-full h-full flex flex-col">
  <!-- 标签栏 -->
  <div class="flex gap-1 bg-surface-100 border-b border-surface-300 px-1">
    {#each tabs as tab (tab.id)}
      <div role="button" tabindex="0"
        class="relative flex items-center gap-2 px-3 py-1.5 rounded-t-md cursor-pointer whitespace-nowrap text-sm
          {activeTab === tab.id ? 'bg-blue-600 text-white border border-b-0 border-surface-300' : 'bg-transparent hover:bg-surface-200'}"
        onclick={() => activeTab = tab.id} onkeydown={(e) => e.key === 'Enter' && (activeTab = tab.id)}
      >
        {tab.title}
        {#if tab.id !== 'home'}
          <button
            onclick={(e) => closeTab(e, tab.id)}
            class="hover:bg-surface-300 rounded-full p-0.5 transition-colors"
          >
            <X size={14} />
          </button>
        {/if}
      </div>
    {/each}
  </div>

  <!-- 内容区域 -->
  <div class="flex-1 bg-white overflow-auto">
    {#each tabs as tab (tab.id)}
      {#if activeTab === tab.id}
        <tab.component />
      {/if}
    {/each}
  </div>
</div>