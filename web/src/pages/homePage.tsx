import { Col, Flex, Row } from "antd"
import { Post } from "../components/posts"
import { getReq } from "../utils/httpClient"
import { UrlConfig } from "../config/urls"
import { useEffect, useState } from "react"
import { AxiosResponse } from "axios"


export const HomePage = () => {
  const config = UrlConfig.fetchBulk
  const [response, setResponse]= useState<{posts:any[]}>()
  const [loading, setLoading]= useState(false)
  const [error, setError]= useState(null)

  useEffect(() => {
    (async ()=>{
      setLoading(true)
      await getReq(config.url, null, {
        params: config.params,
      }).then((res) => {
        console.log(res)

      setResponse((res as AxiosResponse).data)
      })
      .catch((err) => {
        window.location.href = "/auth"
        console.log(err)
        setError(err)
      }).finally(() => {
        setLoading(false)
      })
    })()

  },[config.params, config.url])

  if(error) {
    return <div>Something went wrong!</div>
  }
  if(loading) {
    return <div style={{height:"100vh", backgroundColor: "white"}}>Loading...</div>
  }


  return response?.posts.length ? (
    <>
    <Row gutter={[100, 20]} >
      <>
      {((response as any).posts).map((post: any) => {
        return(
        <Col xs={24} sm={12} md={12} lg={8} xl={8}>
        <Post data={post}/>
        </Col>
      )})}
      </>
    </Row>
    </>
  ):(
    <>
      <Flex justify="center"> No posts Found </Flex>
      <Flex justify="center">
        <div>Want to contribute?</div>
        <div>&nbsp; Go to your Profile and create one!</div>
      </Flex>

    </>
  )
}