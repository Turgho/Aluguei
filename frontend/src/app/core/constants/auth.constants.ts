export const USER_STORAGE_KEY = 'auth_user';

/**
 * Marca requisição já reenviada após refresh.
 * Se voltar 401 com esse header, o interceptor não tenta refresh de novo (evita loop).
 */
export const AUTH_RETRY_HEADER = 'X-Auth-Retry';

/**
 * 401 nessas rotas não dispara refresh no interceptor.
 * /auth/me → só AuthService.fetchMe() renova (evita loop infinito no guard).
 */
export const SKIP_AUTH_REFRESH_PATTERNS = [
  /\/auth\/refresh$/,
  /\/auth\/login$/,
  /\/auth\/logout$/,
  /\/auth\/me$/,
];

/** Rotas de auth que não devem exibir loading global nem simular delay */
export const SKIP_LOADING_PATTERNS = [
  /\/auth\/login$/,
  /\/auth\/register$/,
  /\/auth\/refresh$/,
  /\/auth\/logout$/,
  /\/auth\/me$/,
];
