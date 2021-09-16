import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ignore } from '@cloudinary/base/qualifiers/rotationMode';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-addcoupons',
  templateUrl: './addcoupons.component.html',
  styleUrls: ['./addcoupons.component.css']
})
export class AddcouponsComponent implements OnInit {
  couponObj = {
    couponCode: '',
    couponPrice: ''
  }


  constructor(public http: HttpClient, private dialogRef: MatDialogRef<AddcouponsComponent>, public utilityService: UtilityService, private snackbar: MatSnackBar) { }

  ngOnInit(): void {
  }

  addCoupon() {
    let that = this;
    console.log("this.couponCode", this.couponObj)
    this.http.post('http://localhost:8080/api/addcoupons/', this.couponObj)
      .subscribe(res => {
        console.log("result of the response", res)
        that.utilityService.snackbarmsg = res
        that.snackbar.openFromComponent(SnackbarComponent, {
          duration: 3000,
        });
      });
    this.dialogRef.close();
  }

}
