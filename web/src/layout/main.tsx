import React, { ReactElement, useState } from 'react';
import {
  FolderOpenFilled,
  SmileFilled,
  HomeFilled
} from '@ant-design/icons';
import Logo from '../assets/logo.jpg';
import { Layout, Menu, Typography, theme } from 'antd';
import { ProfilePage } from '../pages/profilePage';
import { HomePage } from '../pages/homePage';

const { Header, Content, Footer, Sider } = Layout;


export const MainLayout = () => {

  const [selectedMenuKey, setSelectedMenuKey] = useState('1');

  const PAGES: Record<string, ReactElement> = {
    "1": <HomePage />,
    '2': <ProfilePage />
  }

  const Headers: Record<string, string> = {
    "1": "All Posts",
    '2': "My Profile",
  }


  const menuItems = [
    {
      key: "1",
      icon: React.createElement(HomeFilled),
      label: "Home",
    },
    {
      key: "2",
      icon: React.createElement(SmileFilled),
      label: "My Profile",
    },
    {
      key: "3",
      icon: React.createElement(FolderOpenFilled),
      label: "My Posts",
    },
  ]

  const {
    token: { colorBgContainer, borderRadiusLG, },
  } = theme.useToken();

  return (
    <Layout >
      <Sider
        breakpoint="xxl"
        collapsedWidth="0"
        width={"27vw"}
        style={{
          zIndex: 5,
          background: "#000",
          height: '100vh',
          position: 'fixed',
          left: 0,
          top: 0,
          bottom: 0,
        }}
      >
        <div className="demo-logo-vertical" style={{ height: "20vh", justifyContent: "center", display: "flex" }}>
          <img height="100vh" src={Logo} alt="logo" />
        </div>
        <Menu style={{ background: "#000" }} theme="dark" mode="inline" defaultSelectedKeys={[selectedMenuKey]} items={menuItems}
          onClick={({ key }) => { setSelectedMenuKey(key) }}
        />
      </Sider>
      <Layout style={{ marginLeft: "0px", width: "100vw", height: "100vh" }}>
        <Header style={{ padding: 0, background: colorBgContainer, display: "flex", justifyContent: "center " }} >
          <Typography.Title level={2} style={{ marginTop: "10px" }} >
            {Headers[selectedMenuKey] ?? ""}
          </Typography.Title>
        </Header>
        <Content style={{ margin: '24px 1px 0', overflow: 'initial', height: "100hv" }}>
          {PAGES[selectedMenuKey] ? PAGES[selectedMenuKey] : (<div
            style={{
              padding: 24,
              textAlign: 'center',
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            <p>long content</p>
            {
              // indicates very long content
              Array.from({ length: 100 }, (_, index) => (
                <React.Fragment key={index}>
                  {index % 20 === 0 && index ? 'more' : '...'}
                  <br />
                </React.Fragment>
              ))
            }
          </div>)}
        </Content>
        <Footer style={{ textAlign: 'center' }}>Gin & ❤️</Footer>
      </Layout>
    </Layout>
  );
}