<script>
  import Login from './pages/Login.svelte'
  import NcnPage from './pages/NcnPage.svelte'
  import Router from 'svelte-spa-router';

  let isLoggedIn = false

  // 监听登录成功事件
  function handleLoginSuccess() {
    isLoggedIn = true
  }
  const routes = {
    '/': NcnPage
  }

  window.addEventListener('login-success', handleLoginSuccess)
</script>

{#if !isLoggedIn}
  <Login />
{:else}
  <Router {routes} />
{/if}
