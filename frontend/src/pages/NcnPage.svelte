<script>
  import { Search } from '../../wailsjs/go/main/Ncn.js'

  // 搜索区域显示/隐藏状态
  let searchSectionVisible = true;

  // 切换搜索区域显示/隐藏
  function toggleSearchSection() {
    searchSectionVisible = !searchSectionVisible;
  }

  // 页面数据
  let searchParams = {
    startTime: '',
    endTime: '',
    ncnNumber: '',
    systemModule: '',
    process: '',
    flow: '',
    riskLevel: '',
    applicant: '',
    status: '',
    mainExceptionCategory: '',
    mainDefect: '',
    responsibleUnit: '',
    ocapProcess: '',
    exceptionBatch: '',
    productModel: '',
    exceptionEquipment: '',
    riskGrade: '',
    responsiblePerson: ''
  }

  // 表格数据
  let tableData = []

  // 加载状态
  let loading = false

  // 错误信息
  let error = ''

  // 计算时效（小时）
  function calculateDuration(receiveTime, finishTime) {
    if (!receiveTime || !finishTime) return ''
    
    // 尝试解析时间格式
    const receiveDate = new Date(receiveTime)
    const finishDate = new Date(finishTime)
    
    if (isNaN(receiveDate.getTime()) || isNaN(finishDate.getTime())) {
      return ''
    }
    
    const durationMs = finishDate.getTime() - receiveDate.getTime()
    const durationHours = (durationMs / (1000 * 60 * 60)).toFixed(2)
    
    return durationHours
  }

  // 格式化时间
  function formatDateTime(dateTimeStr) {
    if (!dateTimeStr) return ''
    
    const date = new Date(dateTimeStr)
    if (isNaN(date.getTime())) return ''
    
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    const hours = String(date.getHours()).padStart(2, '0')
    const minutes = String(date.getMinutes()).padStart(2, '0')
    const seconds = String(date.getSeconds()).padStart(2, '0')
    
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
  }

  // 处理搜索按钮点击
  async function handleSearch() {
    console.log('搜索参数:', searchParams)
    
    // 重置状态
    tableData = []
    error = ''
    loading = true
    
    try {
      // 调用后端搜索API
      const result = await Search(searchParams)
      console.log('搜索结果:', result)
      
      // 处理搜索结果，映射到表格字段
      tableData = result.map(item => {
        return {
          ncnNumber: item.TASK_CODE || '',
          riskLevel: item.RISK_DIMENSION || '',
          mainException: item.MAIN_EXC_TYPE_CODE || '',
          mainDefect: item.MAIN_DEFECT_CODE || '',
          secondaryException: item.SUB_EXC_TYPE_CODE || '',
          secondaryDefect: item.SUB_DEFECT_CODE || '',
          processDescription: item.PRC_DIS || '',
          flowName: item.STAGEDESC || '',
          riskGrade: item.RISK_LEVEL || '',
          applicant: item.CREATED_BY || '',
          sourceCode: item.SOURCE_CODE || '',
          responsibleUnit: item.RESP_DEPT || item.CAR_RESP_DEPT || '',
          responsiblePerson: item.RESP_PERSON || item.CAR_RESP_PERSON || '',
          systemModule: item.SYSTEM_MODULE || '',
          monthlyCount: item.MONCOUNT || '0',
          status: item.STATUS || '',
          auditRemark: item.AUDIT_REMARK || '',
          duration: calculateDuration(item.RECEIVE_TIME, item.FINISH_TIME),
          updateTime: formatDateTime(item.UPDATED),
          createWay: item.TRIGGER_WAY || ''
        }
      })
    } catch (err) {
      console.error('搜索失败:', err)
      error = '搜索失败：' + err
    } finally {
      loading = false
    }
  }

  // 处理重置按钮点击
  function handleReset() {
    Object.keys(searchParams).forEach(key => {
      searchParams[key] = ''
    })
  }

  // 处理新增按钮点击
  function handleAdd() {
    console.log('新增')
  }
  // 处理查看按钮点击
  function handleView() {
    console.log('查看')
  }
  
  // 处理录入按钮点击
  function handleEdit() {
    console.log('录入')
  }
</script>

<div class="min-h-screen bg-gray-50 p-6">
  <!-- 顶部按钮区域 -->
  <div class="flex flex-wrap gap-3 mb-6">
    <button 
      class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors shadow-sm"
      on:click={handleAdd}
    >
      新增
    </button>
    <button 
      class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors shadow-sm"
      on:click={handleEdit}
    >
      录入
    </button>
    <button 
      class="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors shadow-sm"
      on:click={handleView}
    >
      查看
    </button>
    <button 
      class="px-4 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors shadow-sm"
      on:click={toggleSearchSection}
    >
      {searchSectionVisible ? '隐藏搜索' : '显示搜索'}
    </button>
  </div>

  <!-- 搜索区域 -->
  {#if searchSectionVisible}
  <div class="bg-white rounded-lg shadow-md p-6 mb-6">
    <h2 class="text-lg font-semibold text-gray-900 mb-4">搜索条件</h2>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      <!-- 时间条件 -->
      <div class="space-y-2">
        <label for="startTime" class="block text-sm font-medium text-gray-700">起始时间</label>
        <input 
          id="startTime"
          type="date" 
          bind:value={searchParams.startTime} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      
      <div class="space-y-2">
        <label for="endTime" class="block text-sm font-medium text-gray-700">结束时间</label>
        <input 
          id="endTime"
          type="date" 
          bind:value={searchParams.endTime} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>

      <!-- 基本信息 -->
      <div class="space-y-2">
        <label for="ncnNumber" class="block text-sm font-medium text-gray-700">NCN编号</label>
        <input 
          id="ncnNumber"
          type="text" 
          bind:value={searchParams.ncnNumber} 
          placeholder="请输入NCN编号"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      
      <div class="space-y-2">
        <label for="exceptionBatch" class="block text-sm font-medium text-gray-700">异常批号</label>
        <input 
          id="exceptionBatch"
          type="text" 
          bind:value={searchParams.exceptionBatch} 
          placeholder="请输入异常批号"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>
      
      <div class="space-y-2">
        <label for="productModel" class="block text-sm font-medium text-gray-700">产品型号</label>
        <input 
          id="productModel"
          type="text" 
          bind:value={searchParams.productModel} 
          placeholder="请输入产品型号"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        />
      </div>

      <!-- 流程信息 -->
      <div class="space-y-2">
        <label for="process" class="block text-sm font-medium text-gray-700">工序</label>
        <select 
          id="process"
          bind:value={searchParams.process} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="成型">成型</option>
          <option value="电镀">电镀</option>
          <option value="阻焊">阻焊</option>
          <option value="图形">图形</option>
          <option value="表面处理">表面处理</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="flow" class="block text-sm font-medium text-gray-700">流程</label>
        <select 
          id="flow"
          bind:value={searchParams.flow} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="成型后检测">成型后检测</option>
          <option value="电镀后研磨">电镀后研磨</option>
          <option value="MSAP减铜">MSAP减铜</option>
          <option value="ABF压膜前处理">ABF压膜前处理</option>
          <option value="成型后水洗">成型后水洗</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="systemModule" class="block text-sm font-medium text-gray-700">系统/模块</label>
        <select 
          id="systemModule"
          bind:value={searchParams.systemModule} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="SPC">SPC</option>
          <option value="IPQC">IPQC</option>
        </select>
      </div>

      <!-- 异常信息 -->
      <div class="space-y-2">
        <label for="riskLevel" class="block text-sm font-medium text-gray-700">风险层面</label>
        <select 
          id="riskLevel"
          bind:value={searchParams.riskLevel} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="产品">产品</option>
          <option value="设备">设备</option>
          <option value="其他">其他</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="riskGrade" class="block text-sm font-medium text-gray-700">风险等级</label>
        <select 
          id="riskGrade"
          bind:value={searchParams.riskGrade} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="A">A</option>
          <option value="B">B</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="mainExceptionCategory" class="block text-sm font-medium text-gray-700">主异常类别</label>
        <select 
          id="mainExceptionCategory"
          bind:value={searchParams.mainExceptionCategory} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="产品缺陷">产品缺陷</option>
          <option value="过程监控">过程监控</option>
          <option value="异常1">异常1</option>
          <option value="异常2">异常2</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="mainDefect" class="block text-sm font-medium text-gray-700">主缺陷</label>
        <select 
          id="mainDefect"
          bind:value={searchParams.mainDefect} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="指控异常">指控异常</option>
          <option value="孔口异色">孔口异色</option>
          <option value="棕化微蚀量">棕化微蚀量</option>
          <option value="Strip板宽">Strip板宽</option>
          <option value="成型前涨缩">成型前涨缩</option>
          <option value="层压白斑">层压白斑</option>
          <option value="孔壁分离">孔壁分离</option>
        </select>
      </div>

      <!-- 人员信息 -->
      <div class="space-y-2">
        <label for="applicant" class="block text-sm font-medium text-gray-700">申请者</label>
        <select 
          id="applicant"
          bind:value={searchParams.applicant} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="王牌">王牌</option>
          <option value="陈玉聪">陈玉聪</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="responsibleUnit" class="block text-sm font-medium text-gray-700">责任单位</label>
        <select 
          id="responsibleUnit"
          bind:value={searchParams.responsibleUnit} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="新创元">新创元</option>
          <option value="其他">其他</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="responsiblePerson" class="block text-sm font-medium text-gray-700">责任人</label>
        <select 
          id="responsiblePerson"
          bind:value={searchParams.responsiblePerson} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="王牌">王牌</option>
          <option value="陈玉聪">陈玉聪</option>
        </select>
      </div>

      <!-- 其他信息 -->
      <div class="space-y-2">
        <label for="status" class="block text-sm font-medium text-gray-700">状态</label>
        <select 
          id="status"
          bind:value={searchParams.status} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="未完成">未完成</option>
          <option value="已完成">已完成</option>
          <option value="审核中">审核中</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="ocapProcess" class="block text-sm font-medium text-gray-700">OCAP流程</label>
        <select 
          id="ocapProcess"
          bind:value={searchParams.ocapProcess} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="流程1">流程1</option>
          <option value="流程2">流程2</option>
        </select>
      </div>
      
      <div class="space-y-2">
        <label for="exceptionEquipment" class="block text-sm font-medium text-gray-700">异常设备</label>
        <select 
          id="exceptionEquipment"
          bind:value={searchParams.exceptionEquipment} 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        >
          <option value="">请选择</option>
          <option value="设备1">设备1</option>
          <option value="设备2">设备2</option>
        </select>
      </div>
    </div>

    <!-- 搜索和重置按钮 -->
    <div class="flex flex-wrap gap-3 mt-6">
      <button 
        class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors shadow-sm"
        on:click={handleSearch}
        disabled={loading}
      >
        {loading ? '搜索中...' : '查询'}
      </button>
      <button 
        class="px-6 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors shadow-sm"
        on:click={handleReset}
      >
        重置
      </button>
    </div>
  </div>
  {/if}

  <!-- 错误信息显示 -->
  {#if error}
  <div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-6">
    {error}
  </div>
  {/if}

  <!-- 表格区域 -->
  <div class="bg-white rounded-lg shadow-md overflow-hidden">
    <div class="overflow-x-auto">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              NCN编号
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              风险层面
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              主异常
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              主缺陷
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              次异常
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              次缺陷
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              工序描述
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              流程名称
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              风险等级
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              申请者
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              来源系统
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              责任单位
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              责任人
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              来源单号
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              本月发生次数
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              状态
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              审核备注
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              时效(小时)
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              更新时间
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              建单方式
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          {#if loading}
            <tr>
              <td colspan={20} class="px-6 py-12 text-center text-gray-500">
                <div class="flex items-center justify-center">
                  <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mr-3"></div>
                  <span>正在加载数据...</span>
                </div>
              </td>
            </tr>
          {:else if tableData.length === 0}
            <tr>
              <td colspan={20} class="px-6 py-12 text-center text-gray-500">
                暂无数据
              </td>
            </tr>
          {:else}
            {#each tableData as row, index}
              <tr class={index % 2 === 0 ? 'bg-white' : 'bg-gray-50'}>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.ncnNumber}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.riskLevel}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.mainException}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.mainDefect}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.secondaryException}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.secondaryDefect}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.processDescription}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.flowName}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.riskGrade}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.applicant}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.systemModule}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.responsibleUnit}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.responsiblePerson}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.sourceCode}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.monthlyCount}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.status}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.auditRemark}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.duration}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.updateTime}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{row.createWay}</td>
              </tr>
            {/each}
          {/if}
        </tbody>
      </table>
    </div>
  </div>
</div>

<style>
  .animate-spin {
    animation: spin 1s linear infinite;
  }
  
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
</style>
