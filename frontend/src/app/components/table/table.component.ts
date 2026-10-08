import { Component, Input } from '@angular/core';
import { NgTemplateOutlet } from '@angular/common';

export interface TableColumn<T = any> {
  key: string;
  label: string;
  class?: string;
  render?: (row: T) => string;
}

@Component({
  selector: 'app-table',
  standalone: true,
  imports: [NgTemplateOutlet],
  template: `
    <div class="overflow-x-auto rounded-lg border border-border">
      <table class="w-full text-sm">
        <thead class="bg-bg-secondary text-text-secondary">
          <tr>
            @for (col of columns; track col.key) {
              <th class="px-4 py-3 text-left font-medium">{{ col.label }}</th>
            }
            @if (actions) {
              <th class="px-4 py-3 text-right font-medium">Ações</th>
            }
          </tr>
        </thead>
        <tbody class="divide-y divide-border">
          @for (row of rows; track trackBy ? trackBy(row) : $index) {
            <tr class="bg-bg-primary hover:bg-bg-secondary transition-colors">
              @for (col of columns; track col.key) {
                <td class="px-4 py-3 text-text-primary" [class]="col.class || ''">
                  {{ col.render ? col.render(row) : row[col.key] }}
                </td>
              }
              @if (actions) {
                <td class="px-4 py-3 text-right">
                  <ng-container *ngTemplateOutlet="actionsTpl; context: { $implicit: row }"></ng-container>
                </td>
              }
            </tr>
          }
        </tbody>
      </table>
    </div>
    <ng-template #actionsTpl let-row><ng-content /></ng-template>
  `,
})
export class TableComponent {
  @Input() columns: TableColumn[] = [];
  @Input() rows: any[] = [];
  @Input() actions = false;
  @Input() trackBy?: (r: any) => any;
}