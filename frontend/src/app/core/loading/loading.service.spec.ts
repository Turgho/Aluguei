import { TestBed } from '@angular/core/testing';
import { LoadingService } from './loading.service';

describe('LoadingService', () => {
  let service: LoadingService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LoadingService);
  });

  it('deve ser criado', () => {
    expect(service).toBeTruthy();
  });

  it('deve exibir loading após show', () => {
    service.show();
    expect(service.isLoading()).toBe(false);
  });

  it('deve ocultar após hide', () => {
    service.show();
    service.hide();
    expect(service.isLoading()).toBe(false);
  });
});
