// @flow
import * as React from 'react'
import {SendForm, RequestForm} from '.'
import * as WalletsGen from '../../actions/wallets-gen'
import {connect, type TypedState} from '../../util/container'

const mapStateToProps = (state: TypedState) => ({
  isRequest: state.wallets.buildingPayment.isRequest,
})

const mapDispatchToProps = (dispatch, {navigateUp}) => ({
  onClose: () => dispatch(WalletsGen.createAbandonPayment()),
})

const mergeProps = ({isRequest}, {onClose}) => ({
  isRequest,
  onClose,
})

const TheBigSplit = ({isRequest, onClose}) =>
  isRequest ? <RequestForm onClose={onClose} /> : <SendForm onClose={onClose} />

export default connect(
  mapStateToProps,
  mapDispatchToProps,
  mergeProps
)(TheBigSplit)
