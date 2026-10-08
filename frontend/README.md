# Frontend — Aluguei (PWA)

Frontend do Aluguei feito em **Angular** (standalone) com **PWA** (Service Worker), **TailwindCSS** e **gráficos** via Chart.js (`ng2-charts`).

---

## Stack

- **Angular**: `frontend/angular.json`
- **PWA**: `frontend/ngsw-config.json` + `frontend/public/manifest.webmanifest`
- **UI**: TailwindCSS
- **Charts**: Chart.js + `ng2-charts`

---

## Requisitos

| Dependência | Para quê |
|------------|----------|
| **Node.js + npm** | Instalar deps e rodar o Angular |
| **Angular CLI** | `ng serve`, `ng build`, `ng lint`, `ng test` |

---

## Configuração (API)

### Proxy (dev)

O projeto possui `frontend/proxy.conf.json` apontando `/api` para o backend:

- `/api` → `http://localhost:3000`

No `angular.json`, o `serve` usa `"proxyConfig": "proxy.conf.json"`.

### environment

O código referencia `environment` via:

```ts
import { environment } from '../../../environments/environment';
```

e o `angular.json` prevê file replacements em produção:

- `src/environments/environment.ts`
- `src/environments/environment.production.ts`

Se esses arquivos ainda não existirem no seu checkout, crie-os (exemplo mínimo):

```ts
export const environment = {
  apiUrl: '/api/v1',
  useMockAuth: false,
};
```

> O `apiUrl` como `'/api/v1'` funciona bem com o proxy `/api` e as rotas do backend (`/api/v1/...`).

---

## Execução (dev)

```bash
make dev-frontend
```

Ou diretamente:

```bash
npm install
npm start
```

O `ng serve` sobe em `http://localhost:4200`.

---

## Rotas (app)

Definidas em `src/app/app.routes.ts`:

- Públicas: `/login`, `/register`
- Protegidas (layout principal): `/dashboard`, `/properties`, `/tenants`, `/contracts`, `/reports`

---

## Autenticação (browser)

O frontend usa cookies HttpOnly do backend e envia requisições com credenciais:

- `AuthService`: `withCredentials: true`
- `authInterceptor`: clona as requisições com `withCredentials: true`

Isso permite que `access_token`/`refresh_token` sejam enviados automaticamente pelo navegador.

---

## Build

```bash
make build-frontend
```

Ou:

```bash
npm run build
```

Em produção, o Service Worker é habilitado (ver `provideServiceWorker` em `src/app/app.config.ts`).

---

## Testes / qualidade

```bash
make lint-frontend
make test-frontend
```

