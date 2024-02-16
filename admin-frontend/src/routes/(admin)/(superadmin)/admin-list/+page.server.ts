import { routeApi } from "@/lib/util";
import type { PageServerLoad } from "./$types";
import type { ApiResponse } from "@/lib/types/apiResponse";
import type { User } from "@/lib/types/modelTypes";

export const load: PageServerLoad = async({fetch, parent, cookies}) => {
  const {user} = await parent()
  const req = await fetch(routeApi("admin/list-admin"))
  const res: ApiResponse<User[]> = await req.json()

  return { users: res.data, user, token: cookies.get("admin-token") }
}
