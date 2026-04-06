<script>
  import Login from "./pages/Login.svelte";
  import Tabman from "./pages/Tabman.svelte";
  import { EventsOn } from "../wailsjs/runtime/runtime.js";

  let isLoggedIn = $state(false);
  // Listen for Login Success Events
  function handleLoginSuccess() {
    isLoggedIn = true;
  }
  // const routes = {
  //   "/": Tabman,
  //   "/login": Login,
  //   "/ncnpage":NcnPage,
  //   "/tabman":Home,
  // };
  // $effect(() => {
  //   if (isLoggedIn) {

  //   } else {
  //     push("/login");
  //   }
  // });
  EventsOn("exit", () => {
      isLoggedIn = false;
      return;
    //window.location.hash = page;
  });
  window.addEventListener("login-success", handleLoginSuccess);
</script>

<!-- <Router {routes} /> -->
{#if !isLoggedIn}
  <Login />
{:else}
  <Tabman />
{/if}
