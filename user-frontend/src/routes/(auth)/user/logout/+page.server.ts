import { session } from '@/lib/stores/session';
import type { PageServerLoadEvent } from './$types';
import { redirect } from '@sveltejs/kit';

export const load = async (event: PageServerLoadEvent) => {
  session.set({ user: null });
  delete event.locals.user
  event.cookies.delete('token', { path: '/' });
  throw redirect(301, '/login');
};
