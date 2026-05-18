import { TestBed } from '@angular/core/testing';
import { Router, UrlTree } from '@angular/router';
import { firstValueFrom, isObservable } from 'rxjs';
import { authGuard, guestGuard } from './auth.guard';
import { AuthService } from '../auth/auth.service';

describe('auth guards', () => {
  let auth: { fetchMe: ReturnType<typeof vi.fn> };
  let router: { createUrlTree: ReturnType<typeof vi.fn> };

  beforeEach(() => {
    auth = { fetchMe: vi.fn() };
    router = {
      createUrlTree: vi.fn((commands: string[]) => commands as unknown as UrlTree),
    };

    TestBed.configureTestingModule({
      providers: [
        { provide: AuthService, useValue: auth },
        { provide: Router, useValue: router },
      ],
    });
  });

  async function runGuard(guard: typeof authGuard): Promise<unknown> {
    const result = TestBed.runInInjectionContext(() => guard({} as never, {} as never));
    if (isObservable(result)) {
      return firstValueFrom(result);
    }
    return result;
  }

  it('authGuard libera com sessão válida', async () => {
    auth.fetchMe.mockResolvedValue(true);
    await expect(runGuard(authGuard)).resolves.toBe(true);
  });

  it('authGuard redireciona para login sem sessão', async () => {
    auth.fetchMe.mockResolvedValue(false);
    await expect(runGuard(authGuard)).resolves.toEqual(['/login']);
  });

  it('guestGuard redireciona para dashboard com sessão', async () => {
    auth.fetchMe.mockResolvedValue(true);
    await expect(runGuard(guestGuard)).resolves.toEqual(['/dashboard']);
  });

  it('guestGuard libera rota pública sem sessão', async () => {
    auth.fetchMe.mockResolvedValue(false);
    await expect(runGuard(guestGuard)).resolves.toBe(true);
  });
});
