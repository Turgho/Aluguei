import { TestBed } from '@angular/core/testing';
import { ThemeService } from './theme.service';
import { clearStorage } from '../utils/storage.util';

describe('ThemeService', () => {
  let service: ThemeService;

  beforeEach(() => {
    clearStorage(localStorage);
    TestBed.configureTestingModule({});
    service = TestBed.inject(ThemeService);
  });

  it('deve ser criado', () => {
    expect(service).toBeTruthy();
  });

  it('deve alternar o tema', () => {
    const initial = service.theme();
    service.toggle();
    expect(service.theme()).not.toBe(initial);
  });
});
