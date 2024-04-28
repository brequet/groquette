<script>
  import Button from "$lib/components/ui/button/button.svelte";
  import locationService from "$lib/services/LocationService";
</script>

<div class="h-full border-r p-2">
  <Button class="w-full" on:click={locationService.goToHomePage}
    >New Thread</Button
  >

  <h2 class="mt-4 text-lg font-semibold">Chats</h2>
  todo
</div>
