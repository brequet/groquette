import { push, pop, replace } from "svelte-spa-router";

class LocationService {
  private static instance: LocationService;

  private constructor() {}

  static getInstance(): LocationService {
    if (!LocationService.instance) {
      LocationService.instance = new LocationService();
    }
    return LocationService.instance;
  }

  goToHomePage(): void {
    push("/");
  }

  goToChatPage(questionMessage: string, chatModel: string): void {
    push(
      `/chat?${this.toUrlParam("q", questionMessage)}&${this.toUrlParam(
        "model",
        chatModel
      )}`
    );
  }

  private toUrlParam(paramName: string, paramContent: string): string {
    return `${paramName}=${encodeURIComponent(paramContent)}`;
  }
}

const locationService = LocationService.getInstance();

export default locationService;
