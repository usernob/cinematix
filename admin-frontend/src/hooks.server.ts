import { PUBLIC_API_BASE_URL } from '$env/static/public';
import { redirect, type Handle, type HandleFetch } from '@sveltejs/kit';

export const handleFetch: HandleFetch = async ({ request, fetch, event }) => {
  if (request.url.startsWith(PUBLIC_API_BASE_URL)) {
    request.headers.set('Accept', 'application/json');
  }
  if (
    request.url.startsWith(PUBLIC_API_BASE_URL + '/admin') &&
    !request.headers.get('Authorization')
  ) {
    const token = event.cookies.get('admin-token');
    if (token) {
      request.headers.set('Authorization', `Bearer ${token}`);
    }
  }

  console.log(`Making request to ${request.url}`);
  console.log(request.headers);
  return fetch(request);
};

const protectedRoutes = ['/'];

const mustUnprotect = ['/login'];

export const handle: Handle = async ({ event, resolve }) => {
  const token = event.cookies.get('admin-token');

  if (token && mustUnprotect.includes(event.url.pathname)) {
    if (event.locals.user?.role == 'superadmin') {
      throw redirect(302, '/dashboard/superadmin');
    } else {
      throw redirect(302, '/dashboard/petugas');
    }
  }

  if (!token && protectedRoutes.includes(event.url.pathname)) {
    throw redirect(302, '/login');
  }

  return resolve(event);
};
