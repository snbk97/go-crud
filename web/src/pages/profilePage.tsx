import { Row, Col, Divider, Typography, Button, Flex, Modal } from "antd"
import { getReq, postReq } from "../utils/httpClient"
import { useEffect, useState } from "react"
import { Form, Input, SubmitButton } from "formik-antd"
import { Formik } from "formik"


interface IUser {
  email: string,
  name: string,
  username: string,
}

export const ProfilePage = () => {
  const [user, setUser] = useState<IUser>()
  const [modalState, setModelState] = useState<boolean>(false)
  const fetchUser = () => {
    getReq('/me', null, {}).then((res) => {
      setUser(res.data)
    }).catch((err) => {
      console.log(err)
      window.location.href = "/auth"
    })
  }

  const LogOut = () => {
    if (!user?.username) {
      console.log('No user, failed to logout')
      return
    }
    postReq('/auth/logout', { username: user?.username }, {}).then((res) => {

      window.location.href = "/auth"
    })
  }

  const FormSubmit = (data: any) => {
    const { title, body, image_url } = data
    postReq('/post/create', { title, body, image_url }, {})
      .then((res) => {
        setModelState(false)
        window.location.href = "/"
      }).catch((err) => {
        console.log(err)
        window.location.href = "/auth"
      })

  }


  const PostForm = () => {
    return (
      <Formik initialValues={{}} onSubmit={(v) => FormSubmit(v)}>
        <Form>
          <Form.Item name="itemtitle" rules={[{ required: true }]}>
            <Input name="title" placeholder="Enter Title" />
          </Form.Item>
          <Form.Item name="itemimageurl" rules={[{ required: true }]}>
            <Input name="image_url" placeholder="Enter Image Url" />
          </Form.Item>
          <Form.Item name="itembody" rules={[{ required: true }]}>
            <Input.TextArea name="body" rows={4} placeholder="Enter post content" />
          </Form.Item>
          <SubmitButton>Submit</SubmitButton>
        </Form>
      </Formik>
    )
  }

  useEffect(() => {
    fetchUser()
  }, [])

  return (
    <>
      <Row>
        <Col span={24}>
          <Flex justify="center" style={{ width: "100%" }}>
            <Typography.Title level={2}>{`Welcome ${user?.name}`}</Typography.Title>
          </Flex>
          <Divider />

        </Col>
      </Row>
      <Row gutter={[16, 16]}>
        <Col span={18} style={{ width: '100vw' }}>
          <Flex justify="center" style={{ width: "100%", textAlign: "end" }}>
            <Row style={{ display: "block", width: "100%" }}>
              <Typography.Text>Username: {user?.username}</Typography.Text>
            </Row>
            <Row style={{ display: "block", width: "100%" }}>
              <Typography.Text>Email: {user?.email}</Typography.Text>
            </Row>
          </Flex>
        </Col>
        <Divider />
        <Flex justify="center" style={{ width: "100%", textAlign: "end" }}>
          <div style={{ width: "20%", display: "flex", justifyContent: "space-between" }}>
            <Button style={{ color: 'white', background: 'indianred' }} onClick={LogOut}>Logout</Button>
            <Button style={{ color: 'white', background: 'indigo' }} onClick={() => setModelState(true)}>Create Post</Button>
          </div>

        </Flex>
      </Row>
      <Modal title="Create Post"
        open={modalState}
        onCancel={() => setModelState(false)}
        okButtonProps={{ style: { display: 'none' } }}
        cancelButtonProps={{ style: { display: 'none' } }}
      >
        <PostForm />
      </Modal>
    </>
  )
}