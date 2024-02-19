import { routeApi } from '@/lib/util';
import type { PageServerLoad } from './$types';
import type { ApiResponse } from '@/lib/types/apiResponse';
import type { Report } from '@/lib/types/modelTypes';

export const load: PageServerLoad = async ({ fetch }) => {
	const response = await fetch(routeApi('admin/report/thisweek'));
	const data: ApiResponse<Report[]> = await response.json();
	return { data: data.data };
};
