<script>
  import Login from "./pages/Login.svelte";
  import NcnPage from "./pages/NcnPage.svelte";
  import Router from "svelte-spa-router";
  import { EventsOn } from "../wailsjs/runtime/runtime.js";
  import { push } from "svelte-spa-router";
  import { onMount } from "svelte";

  let isLoggedIn = $state(false);

  // Listen for Login Success Events
  function handleLoginSuccess() {
    isLoggedIn = true;
  }
  const routes = {
    "/": NcnPage,
    "/login": Login,
  };
  $effect(() => {
    if (isLoggedIn) {
      push("/"); 
    } else {
      push("/login");
    }
  });
  onMount(() => {
      EventsOn("navigate", (page) => {
        if (page === "/login") {
          isLoggedIn = false;
        }
      });
      window.addEventListener("login-success", handleLoginSuccess);
  });
</script>

<Router {routes} />
