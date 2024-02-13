const BACKEND_URI = "http://localhost:5050";

// Auth Endpoints
export const LOGIN_ENDPOINT = BACKEND_URI + "/login";
export const REGISTER_ENDPOINT = BACKEND_URI + "/register";
export const LOGOUT_ENDPOINT = BACKEND_URI + "/logout";
export const AUTH_USER_ENDPOINT = BACKEND_URI + "/getAuthUserID";

// Blob Storage Endpoints
export const GET_FILES_ENDPOINT = BACKEND_URI + "/getFiles";
export const UPLOAD_FILES_ENDPOINT = BACKEND_URI + "/upload";
export const DELETE_FILES_ENDPOINT = BACKEND_URI + "/delete";

// Post Endpoints
export const POSTS_ENDPOINT = BACKEND_URI + "/posts";
export const CREATE_REPO_ENDPOINT = BACKEND_URI + "/post/create";
export const READ_REPO_ENDPOINT = BACKEND_URI + "/post/read/";
export const UPDATE_REPO_ENDPOINT = BACKEND_URI + "/post/update/";
export const DELETE_REPO_ENDPOINT = BACKEND_URI + "/post/delete/";
