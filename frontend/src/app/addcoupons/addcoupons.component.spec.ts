import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddcouponsComponent } from './addcoupons.component';

describe('AddcouponsComponent', () => {
  let component: AddcouponsComponent;
  let fixture: ComponentFixture<AddcouponsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddcouponsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddcouponsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
