import type { HistoryEntry } from "$lib/entity/history";

class HistoryService {
  private static instance: HistoryService;
  private db: IDBDatabase | null = null;

  private constructor() {
    this.initDB();
  }

  static getInstance(): HistoryService {
    if (!HistoryService.instance) {
      HistoryService.instance = new HistoryService();
    }
    return HistoryService.instance;
  }

  private async initDB() {
    return new Promise<void>((resolve, reject) => {
      const request = indexedDB.open("HistoryDB", 1);

      request.onupgradeneeded = (event) => {
        const db = (event.target as IDBOpenDBRequest).result;
        if (!db.objectStoreNames.contains("history")) {
          db.createObjectStore("history", { keyPath: "id" });
        }
      };

      request.onsuccess = (event) => {
        this.db = (event.target as IDBOpenDBRequest).result;
        resolve();
      };

      request.onerror = (event) => {
        reject((event.target as IDBOpenDBRequest).error);
      };
    });
  }

  async addEntry(id: string, messageContent: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      if (!this.db) {
        reject(new Error("Database not initialized"));
        return;
      }

      const transaction = this.db.transaction(["history"], "readwrite");
      const store = transaction.objectStore("history");
      const entry = {
        id,
        message: messageContent,
        date: new Date(),
      };

      const request = store.add(entry);

      request.onsuccess = () => resolve();
      request.onerror = () => reject(request.error);
    });
  }

  async updateEntry(id: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      if (!this.db) {
        reject(new Error("Database not initialized"));
        return;
      }

      const transaction = this.db.transaction(["history"], "readwrite");
      const store = transaction.objectStore("history");

      const getRequest = store.get(id);
      getRequest.onsuccess = () => {
        if (getRequest.result) {
          getRequest.result.date = new Date();
          const updateRequest = store.put(getRequest.result);
          updateRequest.onsuccess = () => resolve();
          updateRequest.onerror = () => reject(updateRequest.error);
        } else {
          reject(new Error("Entry not found"));
        }
      };
      getRequest.onerror = () => reject(getRequest.error);
    });
  }

  async getAllEntriesSortedByDate(): Promise<HistoryEntry[]> {
    return new Promise<any[]>((resolve, reject) => {
      if (!this.db) {
        reject(new Error("Database not initialized"));
        return;
      }

      const transaction = this.db.transaction(["history"], "readonly");
      const objectStore = transaction.objectStore("history");
      const request = objectStore.getAll();

      request.onsuccess = () => {
        const entries = request.result;
        // Sort entries by date in descending order
        entries.sort((a, b) => b.date - a.date);
        resolve(entries);
      };

      request.onerror = (event) => {
        reject((event.target as IDBRequest).error);
      };
    });
  }
}

const historyService = HistoryService.getInstance();

export default historyService;
