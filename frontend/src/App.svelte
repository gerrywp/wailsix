<script>
  import Login from "./pages/Login.svelte";
  import Tabman from "./pages/Tabman.svelte";
  import { EventsOn } from "../wailsjs/runtime/runtime.js";
  import { onMount } from "svelte";

  let isLoggedIn = $state(false);
  function logout() {
    isLoggedIn = false;
  }
  // Listen for Login Success Events
  function handleLoginSuccess() {
    isLoggedIn = true;
  }
  onMount(() => {
    window.addEventListener("login-success", handleLoginSuccess);
    return () => {
      window.removeEventListener("login-success", handleLoginSuccess);
    };
  });
</script>

<!-- <Router {routes} /> -->
{#if !isLoggedIn}
  <Login />
{:else}
  <Tabman {logout} />
{/if}
