import locationService from "./LocationService";

class ChatService {
  enpointUrl = "http://127.0.0.1:8080/v1/chat";

  private static instance: ChatService;

  private constructor() {}

  static getInstance(): ChatService {
    if (!ChatService.instance) {
      ChatService.instance = new ChatService();
    }
    return ChatService.instance;
  }

  createNewThread(message: string, model: string) {
    locationService.
  }

  async sendMessage(
    content: string,
    chatModel: string
  ): Promise<{ response: string } | null> {
    const uuid = this.uuidv4();

    const message = {
      content,
      chatModel,
    };

    // Store the UUID, message content, and current date in local storage
    localStorage.setItem(
      uuid,
      JSON.stringify({
        content,
        date: new Date().toISOString(),
      })
    );

    try {
      const response = await fetch(this.enpointUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(message),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      const responseData: { response: string } = await response.json();
      return responseData;
    } catch (error) {
      console.error("There was a problem with the fetch operation:", error);
    }

    return null;
  }

  private uuidv4() {
    return "10000000-1000-4000-8000-100000000000".replace(/[018]/g, (c) =>
      (
        +c ^
        (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (+c / 4)))
      ).toString(16)
    );
  }
}

const chatService = ChatService.getInstance();

export default chatService;
