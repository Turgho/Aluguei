export const USER_STORAGE_KEY = 'auth_user';

/** Rotas de auth que não devem exibir loading global nem simular delay */
export const AUTH_ROUTE_PATTERN = /\/auth\/(login|logout|refresh|me|register)/;
