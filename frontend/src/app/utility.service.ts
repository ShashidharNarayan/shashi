import { Injectable } from '@angular/core';
import { DomSanitizer, SafeResourceUrl, SafeUrl} from '@angular/platform-browser';

@Injectable({
  providedIn: 'root'
})
export class UtilityService {
  theatres: any
  movies: any
  recentMovies: any
  snackbarmsg: any
  loggedInObj: any
  selectedMovie: any
  moviesCount: any
  getAllMovies: any
  seats: any

  constructor(private sanitizer : DomSanitizer) { }

  getSanitizedURL(url: any) { 
    let safeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(url);
    return safeUrl
  }

}
