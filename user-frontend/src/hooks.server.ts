import { PUBLIC_API_BASE_URL } from '$env/static/public';
import { redirect, type Handle, type HandleFetch } from '@sveltejs/kit';

export const handleFetch: HandleFetch = async ({ request, fetch, event }) => {
  if (request.url.startsWith(PUBLIC_API_BASE_URL)) {
    request.headers.set('Accept', 'application/json');
  }
  if (request.url.startsWith(PUBLIC_API_BASE_URL + "/user")  && !request.headers.get('Authorization')) {
    const token = event.cookies.get('token');
    if (token) {
      request.headers.set('Authorization', `Bearer ${token}`);
    }
  }
  
  console.log(`Making request to ${request.url}`);
  console.log({request})
  return fetch(request);
};

const protectedRoutes = ['/user'];

const mustUnprotect = ['/login'];

export const handle: Handle = async ({ event, resolve }) => {
  const token = event.cookies.get('token');

  if (!token && protectedRoutes.includes(event.url.pathname)) {
    throw redirect(302, '/login');
  }

  if (token && mustUnprotect.includes(event.url.pathname)) {
    throw redirect(302, '/user');
  }

  return resolve(event);
};
