import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { AuthService } from './auth.service';
import { environment } from '../../../environments/environment';
import { clearStorage } from '../utils/storage.util';

describe('AuthService', () => {
  let service: AuthService;
  let http: HttpTestingController;

  beforeEach(() => {
    clearStorage(localStorage);
    clearStorage(sessionStorage);
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting()],
    });
    service = TestBed.inject(AuthService);
    http = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    http.verify();
  });

  it('deve ser criado', () => {
    expect(service).toBeTruthy();
  });

  it('deve fazer login em modo mock', async () => {
    const original = environment.useMockAuth;
    (environment as { useMockAuth: boolean }).useMockAuth = true;

    await service.login({
      email: 'teste@email.com',
      password: '12345678',
      rememberMe: true,
    });

    expect(service.isLoggedIn()).toBe(true);
    expect(service.user()?.email).toBe('teste@email.com');
    expect(service.user()?.email).toBe('teste@email.com');

    (environment as { useMockAuth: boolean }).useMockAuth = original;
  });

  it('deve validar sessão via fetchMe em modo mock', async () => {
    const original = environment.useMockAuth;
    (environment as { useMockAuth: boolean }).useMockAuth = true;

    await service.login({
      email: 'ana@email.com',
      password: '12345678',
      rememberMe: false,
    });

    const ok = await service.fetchMe();
    expect(ok).toBe(true);
    expect(service.user()?.name).toBe('Usuário Demo');

    (environment as { useMockAuth: boolean }).useMockAuth = original;
  });
});
