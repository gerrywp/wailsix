<script>
  let { navigate } = $props();
</script>

<div class="h-full">
  <div class="max-w-4xl mx-auto">
    <h1 class="text-3xl font-bold text-center text-indigo-800 mb-8">
      欢迎使用单据运维系统
    </h1>
    <div class="grid grid-cols-12 gap-6">
      <div
        class="col-span-12 md:col-span-4 flex flex-col items-center justify-center bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition-shadow duration-300"
      >
        <button
          class="w-full py-3 px-4 bg-green-600 text-white rounded-lg font-medium hover:bg-green-700 transition-colors duration-300"
          onclick={() => navigate("ncn")}>Ncn单据运维</button
        >
      </div>
      <div
        class="col-span-12 md:col-span-4 flex flex-col items-center justify-center bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition-shadow duration-300"
      >
        <button
          class="w-full py-3 px-4 bg-purple-600 text-white rounded-lg font-medium hover:bg-purple-700 transition-colors duration-300"
          onclick={() => navigate("ipqc")}>IPQC单据运维</button
        >
      </div>
    </div>
  </div>
</div>
