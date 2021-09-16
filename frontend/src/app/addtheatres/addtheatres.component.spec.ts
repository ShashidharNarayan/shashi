import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddtheatresComponent } from './addtheatres.component';

describe('AddtheatresComponent', () => {
  let component: AddtheatresComponent;
  let fixture: ComponentFixture<AddtheatresComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddtheatresComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddtheatresComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
