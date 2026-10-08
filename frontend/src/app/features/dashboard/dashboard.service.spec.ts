import { TestBed } from '@angular/core/testing';
import { firstValueFrom } from 'rxjs';
import { DashboardService } from './dashboard.service';

describe('DashboardService', () => {
  let service: DashboardService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DashboardService);
  });

  it('deve retornar dados mock', async () => {
    const data = await firstValueFrom(service.load());
    expect(data.summaryCards.length).toBe(4);
    expect(data.properties.length).toBeGreaterThan(0);
  });
});
