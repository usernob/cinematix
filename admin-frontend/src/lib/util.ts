import { PUBLIC_API_BASE_URL } from '$env/static/public';

export const routeApi = (endpoint: string) => {
	return `${PUBLIC_API_BASE_URL}/${endpoint}`;
};
