import React from "react"
import jdyStyles from "./container.module.css"
import OverLap from "./overlap"

// TAB button 组件
class FloatButton extends React.Component {

    constructor(props){
        super(props);
        this.onMouseOver = this.onMouseOver.bind(this);
        this.onMouseOut = this.onMouseOut.bind(this);

        // 设置默认的over状态
        this.state = {
            isMouseOn: false
        };

    }

    // 监听hover on
    onMouseOver =()=>{
        this.setState({ isMouseOn: true });
    }

    // 监听hover out
    onMouseOut =()=>{
        this.setState({ isMouseOn: false });
    }

    render() {
        const {isTop,icon,onClick,children,leftWindowContent,leftWindowTopDistance} = this.props;
        let leftWindow = null;

        // 获取mouse over状态
        const { isMouseOn } = this.state;

        if(leftWindowContent && isMouseOn){
            leftWindow = <OverLap topDistance={leftWindowTopDistance}>
                {leftWindowContent}
            </OverLap> ;
        }

        return (
            <div className={jdyStyles.side_float_btn} onClick={onClick} onKeyDown={onClick} role="button" tabIndex="0" onFocus={()=>0} onBlur={()=>0} onMouseOver={this.onMouseOver} onMouseOut={this.onMouseOut} >
                <img style={{ margin: (isTop?'12px 0 6px':'8px 0 6px') }}  src={icon} alt="suspend-top.png"></img>
                { children }

                {leftWindow}

            </div>
        )
    }
}

export default FloatButton;