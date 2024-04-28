import About from "./routes/About.svelte";
import Home from "./routes/Home.svelte";
import Chat from "./routes/chat/Chat.svelte";
import NotFound from "./routes/NotFound.svelte";

export const routes = {
  "/": Home,
  "/about": About,
  "/chat": Chat,
  "*": NotFound,
};
