interface IUrlConfig {
  [key:string]:{
  url: string;
  method: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  params: Record<string, unknown>;
  }

}
export const UrlConfig:IUrlConfig= {
  "fetchBulk" : {
    url : "/post/fetchBulk",
    method: "GET",
    params : {
      "limit" : 10,
      "offset" : 0
    }
  }
}