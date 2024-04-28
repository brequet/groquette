<script lang="ts">
  import { AutoResizeTextarea } from "$lib/components/ui/autoresizetextarea";
  import Button from "$lib/components/ui/button/button.svelte";
  import { ArrowRight } from "lucide-svelte";
  import * as Select from "$lib/components/ui/select/index.js";
  import historyService from "$lib/services/HistoryService";
  import locationService from "$lib/services/LocationService";

  // TODO: in onMount, call API to load these
  const chatModels = [
    { value: "llama3-8b-8192", label: "llama3-8b-8192" },
    { value: "llama2-70b-4096", label: "llama2-70b-4096" },
    { value: "mixtral-8x7b-32768", label: "mixtral-8x7b-32768" },
    { value: "gemma-7b-it", label: "gemma-7b-it" },
    { value: "llama3-70b-8192", label: "llama3-70b-8192" },
  ];

  let userMessageInput = "";
  let selectedChatModel = chatModels[0];

  function createThread() {
    console.log(selectedChatModel, userMessageInput);
    locationService.goToChatPage(userMessageInput, selectedChatModel.value);
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.ctrlKey && event.key === "Enter") {
      event.preventDefault();
      createThread();
    }
  }
</script>

<div class="flex flex-col flex-grow h-full items-center">
  <h1 class="text-4xl font-bold mb-8">Groquette</h1>

  <div
    class="flex flex-col w-full max-w-md p-2 mb-4 border border-gray-300 rounded-md"
  >
    <AutoResizeTextarea
      class="w-full resize-none focus-visible:outline-none m-1"
      bind:value={userMessageInput}
      minRows={3}
      maxRows={20}
      on:keydown={handleKeydown}
      placeholder="What is troubling you ?"
    ></AutoResizeTextarea>
    <div class="flex justify-between">
      <Select.Root portal={null} bind:selected={selectedChatModel}>
        <Select.Trigger class="w-[180px]">
          <Select.Value
            placeholder="Select a fruit"
            on:change={() => console.log("change")}
          />
        </Select.Trigger>
        <Select.Content>
          <Select.Group>
            {#each chatModels as chatModel}
              <Select.Item value={chatModel.value} label={chatModel.label}>
                {chatModel.label}
              </Select.Item>
            {/each}
          </Select.Group>
        </Select.Content>
        <Select.Input name="favoriteFruit" />
      </Select.Root>

      <Button variant="ghost" size="icon" on:click={createThread}>
        <ArrowRight class="h-4 w-4" />
      </Button>
    </div>
  </div>
</div>
