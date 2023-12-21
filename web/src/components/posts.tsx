import { Card } from "antd"
const { Meta } = Card;


export interface IPost {
  ID:        number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: null;
  title:     string;
  body:      string;
  image_url: string;
  slug:      string;
  username:  string;
  user:      IUser;
}

export interface IUser {
  ID:        number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: null;
  name:      string;
  username:  string;
  email:     string;
}


export const Post = ({data}:{data: IPost}) => {
const getMeta = () => {
  return data && ( data.body|| "").substring(0, 50) + "..."
}
  return data ? <Card
  hoverable
  style={{ width: 400}}
  cover={<img width="50%" style={{backgroundSize:"contain"}} alt="example" src={data?.image_url} />}
>
  <Meta title={data?.title} description={getMeta()} />
</Card> : <></>
}