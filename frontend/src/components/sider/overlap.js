
import React from "react"
import jdyStyles from "./container.module.css"

// OverLap 组件
class OverLap extends React.Component {
    render() {
        const {children,topDistance} = this.props
        return (
            <div className={jdyStyles.air_bubble} style={{top:topDistance+'px'}}>
                {children}
            </div>
        )
    }
}

export default OverLap;