/** Acesso seguro ao storage (funciona em testes sem localStorage) */
export function getStorageItem(storage: Storage | undefined, key: string): string | null {
  if (!storage) return null;
  try {
    return storage.getItem(key);
  } catch {
    return null;
  }
}

export function setStorageItem(storage: Storage | undefined, key: string, value: string): void {
  if (!storage) return;
  try {
    storage.setItem(key, value);
  } catch {
    // ambiente de teste ou modo privado
  }
}

export function removeStorageItem(storage: Storage | undefined, key: string): void {
  if (!storage) return;
  try {
    storage.removeItem(key);
  } catch {
    // ambiente de teste
  }
}

export function clearStorage(storage: Storage | undefined): void {
  if (!storage) return;
  try {
    storage.clear();
  } catch {
    // ambiente de teste
  }
}
