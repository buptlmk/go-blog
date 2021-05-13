import React from "react"
import { css } from "@emotion/core"
import FloatButton from "./float-icon-text-button"

// 回到顶部图标
const imageSuspendTop = require("../images/suspend-top.png");
// 客服图标
const imageSuspendPhone = require("../images/suspend-phone.png");
// 公众号图标
const imageSuspendWx = require("../images/suspend-wx.png");
// 申请使用
const imageSuspendUse = require("../images/suspend-use.png");

// 弹框 客服图标
const imageSuspendService = require("../images/suspend-service.png");
// 弹框 二维码图标
const imageSuspendCode = require("../images/suspend-code.png");


// 定义sidebar组件
export class SideBar extends React.Component {

    constructor(props){
        super(props);

        this.click2Top = this.click2Top.bind(this);
        this.click2PhoneService = this.click2PhoneService.bind(this);
        this.click2WechatAccount = this.click2WechatAccount.bind(this);
        this.click2ApplyUsing = this.click2ApplyUsing.bind(this);

    }

    // 滚到页面顶部
    click2Top =()=>{
        document.body.scrollIntoView(true);//为ture返回顶部，false为底部
    }

    // 点击客服
    click2PhoneService =()=>{
        console.log('111111111111111111  点击客服');
    }

    // 微信公众号
    click2WechatAccount =()=>{
        console.log('111111111111111111 点击微信公众号');
    }

    // 申请使用
    click2ApplyUsing =()=>{
        console.log('111111111111111111 点击申请使用');
    }

    render() {
        return (
            <div css={css`position: fixed;top: 500px;z-index: 10000;right: 68px;`}>

                {/* 顶部 */}
                <FloatButton isTop={true} icon={imageSuspendTop} onClick={()=>this.click2Top()}>顶部</FloatButton>

                {/* 客服 */}
                <FloatButton icon={imageSuspendPhone} onClick={this.click2PhoneService} leftWindowContent={<div css={css`display: flex;align-items: center;flex-direction: column;`}>

                    <div css={css`display: flex;flex-direction: row;`}>
                        <img css={css`margin:0 auto;`} src={imageSuspendService} alt="客服电话"></img>
                        <div css={css`margin-left:6px;color: #333;font-size: 16px;`}>客服电话</div>
                    </div>

                    <div css={css`margin-left:6px;color: #0084ff;font-size: 22px;`}>400-777-585</div>

                </div>} leftWindowTopDistance={84}>客服</FloatButton>

                {/* 公众号 */}
                <FloatButton icon={imageSuspendWx} onClick={this.click2WechatAccount} leftWindowContent={<div css={css`display: flex;align-items: center;flex-direction: column;`}>
                    <div css={css`display: flex;flex-direction: column;`}>
                        <img css={css`margin:0 auto;`} src={imageSuspendCode} alt="客服电话"></img>
                        <div css={css`margin-top:8px;color: #333;font-size: 14px;`}>金斗云智能管理平台订阅号</div>
                    </div>
                    <div css={css`display: flex;flex-direction: column;margin-top:24px;`}>
                        <img css={css`margin:0 auto;`} src={imageSuspendCode} alt="客服电话"></img>
                        <div css={css`margin-top:8px;color: #333;font-size: 14px;`}>金斗云智能管理平台服务号</div>
                    </div>
                </div>} leftWindowTopDistance={0}>公众号</FloatButton>

                {/* 申请使用 */}
                <FloatButton icon={imageSuspendUse} onClick={this.click2ApplyUsing} >申请使用</FloatButton>

            </div>

        )
    }
}
// export default SideBar;