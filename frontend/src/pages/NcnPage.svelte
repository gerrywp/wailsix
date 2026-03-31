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

<div class="main-page">
  <!-- 顶部按钮区域 -->
  <div class="top-buttons">
    <button class="btn-add" on:click={handleAdd}>新增</button>
    <button class="btn-add" on:click={handleEdit}>录入</button>
    <button class="btn-view" on:click={handleView}>查看</button>
    <button class="btn-toggle-search" on:click={toggleSearchSection}>
      {searchSectionVisible ? '隐藏搜索' : '显示搜索'}
    </button>
  </div>

  <!-- 搜索区域 -->
  {#if searchSectionVisible}
  <div class="search-section">
    <div class="search-row">
      <div class="search-item">
        <label>起始时间:</label>
        <input type="date" bind:value={searchParams.startTime} />
      </div>
      <div class="search-item">
        <label>工序:</label>
        <select bind:value={searchParams.process}>
          <option value="">请选择</option>
          <option value="成型">成型</option>
          <option value="电镀">电镀</option>
          <option value="阻焊">阻焊</option>
          <option value="图形">图形</option>
          <option value="表面处理">表面处理</option>
        </select>
      </div>
      <div class="search-item">
        <label>流程:</label>
        <select bind:value={searchParams.flow}>
          <option value="">请选择</option>
          <option value="成型后检测">成型后检测</option>
          <option value="电镀后研磨">电镀后研磨</option>
          <option value="MSAP减铜">MSAP减铜</option>
          <option value="ABF压膜前处理">ABF压膜前处理</option>
          <option value="成型后水洗">成型后水洗</option>
        </select>
      </div>
      <div class="search-item">
        <label>风险层面:</label>
        <select bind:value={searchParams.riskLevel}>
          <option value="">请选择</option>
          <option value="产品">产品</option>
          <option value="设备">设备</option>
          <option value="其他">其他</option>
        </select>
      </div>
      <div class="search-item">
        <label>申请者:</label>
        <select bind:value={searchParams.applicant}>
          <option value="">请选择</option>
          <option value="王牌">王牌</option>
          <option value="陈玉聪">陈玉聪</option>
        </select>
      </div>
    </div>

    <div class="search-row">
      <div class="search-item">
        <label>结束时间:</label>
        <input type="date" bind:value={searchParams.endTime} />
      </div>
      <div class="search-item">
        <label>系统/模块:</label>
        <select bind:value={searchParams.systemModule}>
          <option value="">请选择</option>
          <option value="SPC">SPC</option>
          <option value="IPQC">IPQC</option>
        </select>
      </div>
      <div class="search-item">
        <label>状态:</label>
        <select bind:value={searchParams.status}>
          <option value="">请选择</option>
          <option value="未完成">未完成</option>
          <option value="已完成">已完成</option>
        </select>
      </div>
      <div class="search-item">
        <label>主异常类别:</label>
        <select bind:value={searchParams.mainExceptionCategory}>
          <option value="">请选择</option>
          <option value="产品缺陷">产品缺陷</option>
          <option value="过程监控">过程监控</option>
          <option value="异常1">异常1</option>
          <option value="异常2">异常2</option>
        </select>
      </div>
      <div class="search-item">
        <label>主缺陷:</label>
        <select bind:value={searchParams.mainDefect}>
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
    </div>

    <div class="search-row">
      <div class="search-item">
        <label>责任单位:</label>
        <select bind:value={searchParams.responsibleUnit}>
          <option value="">请选择</option>
          <option value="新创元">新创元</option>
          <option value="其他">其他</option>
        </select>
      </div>
      <div class="search-item">
        <label>OCAP流程:</label>
        <select bind:value={searchParams.ocapProcess}>
          <option value="">请选择</option>
          <option value="流程1">流程1</option>
          <option value="流程2">流程2</option>
        </select>
      </div>
      <div class="search-item">
        <label>异常设备:</label>
        <select bind:value={searchParams.exceptionEquipment}>
          <option value="">请选择</option>
          <option value="设备1">设备1</option>
          <option value="设备2">设备2</option>
        </select>
      </div>
      <div class="search-item">
        <label>风险等级:</label>
        <select bind:value={searchParams.riskGrade}>
          <option value="">请选择</option>
          <option value="A">A</option>
          <option value="B">B</option>
        </select>
      </div>
      <div class="search-item">
        <label>责任人:</label>
        <select bind:value={searchParams.responsiblePerson}>
          <option value="">请选择</option>
          <option value="王牌">王牌</option>
          <option value="陈玉聪">陈玉聪</option>
        </select>
      </div>
    </div>

    <div class="search-row">
      <div class="search-item">
        <label>NCN编号:</label>
        <input type="text" bind:value={searchParams.ncnNumber} />
      </div>
      <div class="search-item">
        <label>异常批号:</label>
        <input type="text" bind:value={searchParams.exceptionBatch} />
      </div>
      <div class="search-item">
        <label>产品型号:</label>
        <input type="text" bind:value={searchParams.productModel} />
      </div>
      <div class="search-item">
        <button class="btn-search" on:click={handleSearch}>查询</button>
      </div>
      <div class="search-item">
        <button class="btn-reset" on:click={handleReset}>重置</button>
      </div>
    </div>
  </div>
  {/if}

  <!-- 表格区域 -->
  <div class="table-section">
    <table>
      <thead>
        <tr>
          <th>NCN编号</th>
          <th>风险层面</th>
          <th>主异常</th>
          <th>主缺陷</th>
          <th>次异常</th>
          <th>次缺陷</th>
          <th>工序描述</th>
          <th>流程名称</th>
          <th>风险等级</th>
          <th>申请者</th>
          <th>来源系统</th>
          <th>责任单位</th>
          <th>责任人</th>
          <th>来源单号</th>
          <th>本月发生次数</th>
          <th>状态</th>
          <th>审核备注</th>
          <th>时效(小时)</th>
          <th>更新时间</th>
          <th>建单方式</th>
        </tr>
      </thead>
      <tbody>
        {#each tableData as row, index}
          <tr class={index % 2 === 0 ? 'even' : 'odd'}>
            <td>{row.ncnNumber}</td>
            <td>{row.riskLevel}</td>
            <td>{row.mainException}</td>
            <td>{row.mainDefect}</td>
            <td>{row.secondaryException}</td>
            <td>{row.secondaryDefect}</td>
            <td>{row.processDescription}</td>
            <td>{row.flowName}</td>
            <td>{row.riskGrade}</td>
            <td>{row.applicant}</td>
            <td>{row.systemModule}</td>
            <td>{row.responsibleUnit}</td>
            <td>{row.responsiblePerson}</td>
            <td>{row.sourceCode}</td>
            <td>{row.monthlyCount}</td>
            <td>{row.status}</td>
            <td>{row.auditRemark}</td>
            <td>{row.duration}</td>
            <td>{row.updateTime}</td>
            <td>{row.createWay}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>

<style>
  .main-page {
    padding: 20px;
    font-family: Arial, sans-serif;
    background-color: #f5f7fa;
    min-height: 100vh;
    box-sizing: border-box;
  }

  /* 顶部按钮样式 */
  .top-buttons {
    margin-bottom: 20px;
    display: flex;
    justify-content: flex-start;
  }

  .btn-add, .btn-view, .btn-search, .btn-reset, .btn-toggle-search {
    padding: 8px 16px;
    margin-right: 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 14px;
    transition: background-color 0.3s;
  }

  .btn-toggle-search {
    background-color: #9C27B0;
    color: white;
  }

  .btn-toggle-search:hover {
    background-color: #7B1FA2;
  }

  .btn-add {
    background-color: #4CAF50;
    color: white;
  }

  .btn-add:hover {
    background-color: #45a049;
  }

  .btn-view {
    background-color: #2196F3;
    color: white;
  }

  .btn-view:hover {
    background-color: #0b7dda;
  }

  .btn-search {
    background-color: #2196F3;
    color: white;
  }

  .btn-search:hover {
    background-color: #0b7dda;
  }

  .btn-reset {
    background-color: #f44336;
    color: white;
  }

  .btn-reset:hover {
    background-color: #da190b;
  }

  /* 搜索区域样式 */
  .search-section {
    background-color: white;
    padding: 15px;
    border-radius: 8px;
    margin-bottom: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    overflow-x: auto;
    white-space: nowrap;
    width: 100%; /* 使搜索区域宽度横铺满屏幕 */
    box-sizing: border-box; /* 确保内边距不会增加总宽度 */
  }

  .search-row {
    display: flex;
    gap: 10px;
    margin-bottom: 10px;
    align-items: center;
  }

  .search-item {
    display: flex;
    align-items: center;
    gap: 6px;
    /* 为每个搜索项设置固定宽度，确保所有项对齐 */
    width: 280px; /* 标签宽度 + 间隙 + 输入框宽度 = 70px + 6px + 200px = 280px */
  }

  /* 调整按钮所在的搜索项，使按钮居中对齐 */
  .search-item:has(.btn-search),
  .search-item:has(.btn-reset) {
    justify-content: flex-start;
    /* 保持与其他搜索项相同的固定宽度，确保对齐 */
    width: 65px;
  }

  .search-item label {
    font-size: 13px;
    color: #333;
    width: 70px; /* 固定标签宽度，确保所有标签对齐 */
    text-align: right;
  }

  .search-item input,
  .search-item select {
    padding: 5px 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 13px;
    width: 200px;
  }

  .search-item input:focus,
  .search-item select:focus {
    outline: none;
    border-color: #2196F3;
  }

  .search-buttons {
    margin-top: 15px;
  }

  /* 表格区域样式 */
  .table-section {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    max-height: 456px; /* 你想要的固定高度，可修改 */
    overflow-y: auto; /* 垂直滚动 */
    overflow-x: auto; /* 列太多时水平滚动 */
  }

  table {
    width: 100%;
    border-collapse: collapse;
    white-space: nowrap; /* 所有文字强制不换行 */
  }
  /* 固定表头核心样式 */
  table thead {
    position: sticky;
    top: 0;
    z-index: 1;
  }
  table th,
  table td {
    padding: 12px;
    text-align: left;
    border-bottom: 1px solid #ddd;
    font-size: 14px;
    color:#333;
    white-space: nowrap; /* 所有文字强制不换行 */
  }

  table th {
    background-color: #f2f2f2;
    font-weight: bold;
    color: #333;
  }

  table tbody tr:hover {
    background-color: #f5f5f5;
  }

  table tbody tr.even {
    background-color: #fafafa;
  }

  /* 响应式设计 */
  @media (max-width: 768px) {
    .search-row {
      flex-wrap: wrap;
    }

    .search-item {
      flex-basis: calc(50% - 10px);
    }

    .table-section {
      font-size: 12px;
    }

    table th,
    table td {
      padding: 8px;
    }
  }
</style>