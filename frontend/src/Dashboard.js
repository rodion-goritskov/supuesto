import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import ScenariosPage from './components/ScenariosPage';
import Menu from './components/Menu';
import { BrowserRouter as Router, Route} from "react-router-dom";
import CreateScenarioPage from "./components/CreateScenarioPage";

const styles = ({
    root: {
        display: 'flex',
    },
});

class Dashboard extends React.Component {
    render() {
        const { classes } = this.props;

        return (
                <Router>
                    <div className={classes.root}>
                        <Menu />
                        <Route path="/" exact component={ScenariosPage} />
                        <Route path="/scenario/create" component={CreateScenarioPage} />
                    </div>
                </Router>
        );
    }
}

Dashboard.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Dashboard);