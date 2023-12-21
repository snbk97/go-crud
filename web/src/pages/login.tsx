import { Card, Col, Flex, Row } from "antd"
import { Formik } from "formik"
import { Form, Input, Radio, SubmitButton } from "formik-antd"
import { useState } from "react"
import { postReq } from "../utils/httpClient"

export const LoginPage = () => {

  interface IAuthData {
    username?: string;
    password?: string;
    email?: string;
    name?: string;
  }
  const Login = (data: IAuthData) => {
    postReq("/auth/login", { ...data }, {})
      .then((res) => {
        console.log("logged in")
        window.location.href = "/"
      }).catch((err) => {
        console.log(err)
      })
  }

  const Signup = (data: IAuthData) => {
    const { username, password, email, name } = data
    postReq("/user", { username, password, email, name }, {})
      .then((res) => {
        console.log("signed up")
        Login({ username, password })
      }).catch((err) => {
        console.log(err)
      })
  }


  const FormSubmit = (data: any) => {
    const { username, password, email, name } = data
    console.log(username, password, email, name)

    const isLogin = (!email && !name) && (username && password)
    const isSignup = (email && name) && (username && password)
    if (isLogin) {
      Login(data)
    }

    if (isSignup) {
      Signup(data)
    }
    return
  }

  const [authFormState, setAuthFormState] = useState<"login" | "signup">("login")

  const AuthForm = () => {
    return (
      <Card>
        <Formik initialValues={{}} onSubmit={(v) => FormSubmit(v)}>
          <Form
            name="basic"
            autoComplete="off"
          >
            <div style={{ padding: "20px 200px" }}>
              <Radio.Group name="authtype" defaultValue={authFormState}>
                <Radio.Button value="login" onClick={() => setAuthFormState("login")}>Login</Radio.Button>
                <Radio.Button value="signup" onClick={() => setAuthFormState("signup")}>Signup</Radio.Button>
              </Radio.Group>
            </div>
            {authFormState === "login" && (<>
              <Form.Item name="itemusername" hasFeedback rules={[{ required: true, message: 'Please enter username!' }]}>
                <Input name="username" placeholder="Enter username" />
              </Form.Item>
              <Form.Item name="itempassword" hasFeedback rules={[{ required: true, message: 'Please enter password' }]}>
                <Input name="password" type="password" placeholder="Enter password" />
              </Form.Item>
              <SubmitButton >Login</SubmitButton>
            </>)}
            {authFormState === "signup" && (<>
              <Form.Item name="itemname" rules={[{ required: true }]}>
                <Input name="name" placeholder="Enter your name" />
              </Form.Item>
              <Form.Item name="itemusername" rules={[{ required: true }]}>
                <Input name="username" placeholder="Enter your username" />
              </Form.Item>
              <Form.Item name="itememail" rules={[{ required: true }]}>
                <Input name="email" type="email" placeholder="Enter email" />
              </Form.Item>
              <Form.Item name="itempassword" rules={[{ required: true }]}>
                <Input name="password" type="password" placeholder="Enter password" />
              </Form.Item>
              <SubmitButton>Signup</SubmitButton>
            </>)}
          </Form>
        </Formik>
      </Card>
    )
  }

  return (
    <Row gutter={[16, 16]} style={{ backgroundColor: "grey", height: "100vh" }}>
      <Col xs={24} sm={24} md={24} lg={18} xl={24} flex={2} >
        <Flex justify="center" style={{ marginTop: "25vh" }}>
          <AuthForm />
        </Flex>

      </Col>
    </Row>
  )

}