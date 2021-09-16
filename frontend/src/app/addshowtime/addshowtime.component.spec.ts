import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddshowtimeComponent } from './addshowtime.component';

describe('AddshowtimeComponent', () => {
  let component: AddshowtimeComponent;
  let fixture: ComponentFixture<AddshowtimeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddshowtimeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddshowtimeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
