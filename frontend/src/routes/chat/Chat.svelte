<script lang="ts">
  import { onMount } from "svelte";
  import { location, querystring } from "svelte-spa-router";
  import { get } from "svelte/store";

  let userMessages: string[] = [];
  let model: string;

  onMount(() => {
    const params = get(querystring)
      ?.split("&")
      .reduce(
        (acc, pair) => {
          const [key, value] = pair.split("=");
          acc[key] = decodeURIComponent(value);
          return acc;
        },
        {} as { [key: string]: string }
      );
    console.log("params", params);
    if (params === undefined) return; // TODO fix
    console.log("params[q]", params["q"]);

    userMessages.push(params["q"]);
    model = params["model"];
  });
</script>

<p>chat page</p>
<p>
  {$location}
</p>

<p class="bg-red-100">{model}</p>
<p>{userMessages[0]}</p>
