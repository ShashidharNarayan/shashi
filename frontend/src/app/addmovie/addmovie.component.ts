import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
// Import the Cloudinary classes.
import { CloudinaryImage } from '@cloudinary/base';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UtilityService } from '../utility.service';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-addmovie',
  templateUrl: './addmovie.component.html',
  styleUrls: ['./addmovie.component.css']
})

export class AddmovieComponent implements OnInit {
  img: CloudinaryImage = new CloudinaryImage;
  data = new FormData();

  movieObj = {
    movieName: '',
    trailerLink: '',
    movieDetails: '',
    screened_at: '',
    file: ''
  }

  constructor(public http: HttpClient, public snackbar: MatSnackBar, public utilityService: UtilityService, private dialogRef: MatDialogRef<AddmovieComponent>) { }

  ngOnInit(): void {
  }

  onChange(event: any) {
    this.movieObj.file = event.target.files[0]
  }

  addMovie(movie: any) {
    // let date = new Date(this.movieObj.screened_at)
    const fd = new FormData()
    fd.append("image", this.movieObj.file)
    fd.append("Content-Type", "multipart/form-data")
    fd.append("movieName", this.movieObj.movieName)
    fd.append("trailerLink", this.movieObj.trailerLink)
    fd.append("movieDetails", this.movieObj.movieDetails)
    // convert object to string then trim it to yyyy-mm-dd
    const stringified = JSON.stringify(this.movieObj.screened_at);
    const date = stringified.substring(1, 11);
    fd.append("screened_at", date)
    console.log("this.movieObj", this.movieObj)
    let that = this;
    this.http.post('http://localhost:8080/api/movie/', fd)
      .subscribe(function (res) {
        console.log("result of movie insertion", res)
        that.utilityService.snackbarmsg = res
        that.snackbar.openFromComponent(SnackbarComponent, {
          duration: 3000,
          data: res
        });
        that.dialogRef.close();
      });
  }
}
