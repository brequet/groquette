<script lang="ts">
  import Router from "svelte-spa-router";
  import { routes } from "./routes";
  import SideBar from "./components/side-bar/SideBar.svelte";
</script>

<main class="grid grid-flow-col min-h-screen">
  <div class="lg:col-span-1 col-span-2 hidden md:block">
    <SideBar />
  </div>
  <div class="lg:col-span-9 col-span-7 flex flex-col">
    <div class="flex flex-row items-center space-x-4 mb-16">
      <a href="#/" class="text-xl text-blue-500 hover:text-blue-700"
        >Hello page</a
      >
      <a href="#/about" class="text-xl text-blue-500 hover:text-blue-700"
        >About page</a
      >
      <a href="#/nothing" class="text-xl text-blue-500 hover:text-blue-700"
        >This page does not exist</a
      >
    </div>

    <div class="flex-grow">
      <Router {routes} class=" h-full " />
    </div>
  </div>
</main>

<!-- <main class="flex flex-col items-center justify-center min-h-screen">
  <h1 class="text-4xl font-bold mb-8">Brequet Template</h1>

  <div class="flex flex-row items-center space-x-4 mb-16">
    <a href="#/" class="text-xl text-blue-500 hover:text-blue-700">Hello page</a
    >
    <a href="#/about" class="text-xl text-blue-500 hover:text-blue-700"
      >About page</a
    >
    <a href="#/nothing" class="text-xl text-blue-500 hover:text-blue-700"
      >This page does not exist</a
    >
  </div>

  <div class="flex-grow flex">
    <Router {routes} class="flex-grow" />
  </div>
</main> -->
