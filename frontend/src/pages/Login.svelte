<script>
  import { Login } from '../../wailsjs/go/main/App.js'

  let username = ''
  let password = ''
  let error = ''

  function handleLogin() {
    error = ''
    if (!username || !password) {
      error = '请输入用户名和密码'
      return
    }

    Login(username, password).then(success => {
      if (success) {
            // 触发登录成功事件
            window.dispatchEvent(new CustomEvent('login-success'))
      } else {
        error = '用户名或密码错误'
      }
    }).catch(err => {
      error = '登录失败：' + err
    })
  }

  function handleKeyPress(event) {
    if (event.key === 'Enter') {
      handleLogin()
    }
  }
</script>

<main class="min-h-screen w-full flex items-center justify-center bg-gradient-to-br from-indigo-600 via-purple-600 to-pink-500">
  <div class="w-full max-w-md px-4">
    <div class="bg-white/10 backdrop-blur-xl rounded-xl shadow-2xl p-8 border border-white/20">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-white mb-2">欢迎登录</h1>
        <p class="text-white/80">请输入您的账号信息</p>
      </div>
      
      {#if error}
        <div class="bg-red-50 border border-red-200 text-red-600 px-4 py-3 rounded-lg mb-6 flex items-start gap-3">
          <svg class="w-5 h-5 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"/>
          </svg>
          <div class="flex-1">
            {error}
          </div>
          <button 
            on:click={() => error = ''}
            class="text-red-400 hover:text-red-600 transition-colors"
          >
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
            </svg>
          </button>
        </div>
      {/if}
      
      <div class="space-y-6">
        <div>
          <label for="username" class="block text-sm font-medium text-white mb-2">
            用户名
          </label>
          <div class="relative">
            <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
              <svg class="w-5 h-5 text-white/70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
              </svg>
            </div>
            <input
              id="username"
              type="text"
              bind:value={username}
              placeholder="请输入用户名"
              autocomplete="off"
              class="w-full pl-12 pr-4 py-3 bg-white/10 backdrop-blur-sm border-2 border-white/30 rounded-lg focus:ring-2 focus:ring-white/50 focus:border-white/50 transition-all duration-200 text-white placeholder:text-white/60"
              on:keypress={handleKeyPress}
              on:input={() => error = ''}
            />
          </div>
        </div>
        
        <div>
          <label for="password" class="block text-sm font-medium text-white mb-2">
            密码
          </label>
          <div class="relative">
            <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
              <svg class="w-5 h-5 text-white/70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
              </svg>
            </div>
            <input
              id="password"
              type="password"
              bind:value={password}
              placeholder="请输入密码"
              class="w-full pl-12 pr-4 py-3 bg-white/10 backdrop-blur-sm border-2 border-white/30 rounded-lg focus:ring-2 focus:ring-white/50 focus:border-white/50 transition-all duration-200 text-white placeholder:text-white/60"
              on:keypress={handleKeyPress}
              on:input={() => error = ''}
            />
          </div>
        </div>
        
        <button
          on:click={handleLogin}
          class="w-full bg-white hover:bg-white/90 text-indigo-600 font-semibold py-3 px-6 rounded-lg transition-all duration-200 shadow-lg hover:shadow-xl transform hover:-translate-y-0.5 focus:outline-none focus:ring-2 focus:ring-white/50 focus:ring-offset-2"
        >
          <div class="flex items-center justify-center gap-2">
            <span>登录</span>
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/>
            </svg>
          </div>
        </button>
      </div>
    </div>
  </div>
</main>