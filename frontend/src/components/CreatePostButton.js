import React, { Component } from "react";
import { Modal, Button, message } from "antd";
import axios from "axios";

import { PostForm } from "./PostForm";
import { BASE_URL, TOKEN_KEY } from "../constants";

class CreatePostButton extends Component {
  state = {
    visible: false,
    confirmLoading: false, //上传图片会花时间，用来confirm正在loading
  };

  showModal = () => {
    this.setState({
      visible: true,
    });
  };

  handleOk = () => {
    this.setState({
      confirmLoading: true,
    });

    // get form data
    this.postForm
      .validateFields()
      .then((form) => {
        //如果用户信息都填好了，调用后端api
        const { description, uploadPost } = form;
        const { type, originFileObj } = uploadPost[0];
        const postType = type.match(/^(image|video)/g)[0];
        if (postType) {
          let formData = new FormData();
          formData.append("message", description);
          formData.append("media_file", originFileObj);

          const opt = {
            method: "POST",
            url: `${BASE_URL}/upload`,
            headers: {
              Authorization: `Bearer ${localStorage.getItem(TOKEN_KEY)}`,
            },
            data: formData,
          };

          axios(opt)
            .then((res) => {
              if (res.status === 200) {
                message.success("The image/video is uploaded!");
                this.postForm.resetFields();
                this.handleCancel();
                this.props.onShowPost(postType); //onshowpost重新调用后端search API,显示上传后的所有图片
              }
            })
            .catch((err) => {
              console.log("Upload image/video failed: ", err.message);
              message.error("Failed to upload image/video!");
            })
            .finally(() => {
              this.setState({ confirmLoading: false });
            });
        }
      })
      .catch((err) => {
        console.log("err validate form -> ", err);
      });
  };

  handleCancel = () => {
    //又不想上传的情况
    console.log("Clicked cancel button");
    this.setState({
      visible: false,
    });
  };

  render() {
    const { visible, confirmLoading } = this.state;
    return (
      <div>
        <Button type="primary" onClick={this.showModal}>
          Create New Post
        </Button>
        <Modal
          title="Create New Post"
          open={visible}
          onOk={this.handleOk}
          okText="Create"
          confirmLoading={confirmLoading}
          onCancel={this.handleCancel}
        >
          <PostForm ref={(refInstance) => (this.postForm = refInstance)} />
        </Modal>
      </div>
    );
  }
}
export default CreatePostButton;
