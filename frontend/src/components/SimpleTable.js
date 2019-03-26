import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import axios from 'axios';

const styles = {
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 700,
    },
};

class ScenarioTable extends React.Component {
    state = {
        scenarios: []
    };

    componentDidMount() {
        axios.get(`http://localhost:4444/v1/scenario`, {
            params: {
                offset: 0,
                max: 2
            }
        }).then(res => {
            const scenarios = res.data;
            this.setState({ scenarios: scenarios })
        })
    }

    render() {
        return (
            <Paper>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Scenario ID</TableCell>
                        <TableCell>Scenario name</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    { this.state.scenarios.map(scenario =>
                        <TableRow key={scenario.uid}>
                            <TableCell component="th" scope="row">
                                {scenario.uid}
                            </TableCell>
                            <TableCell component="th" scope="row">
                                {scenario.name}
                            </TableCell>
                        </TableRow>
                    )}
                </TableBody>
            </Table>
            </Paper>
        );
    }
}

export default withStyles(styles)(ScenarioTable);