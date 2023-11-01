import React from 'react';
import { AppConstants } from '../constants/app-constants';
import { useEffect } from 'react';

interface DiagnosticsRespoonse {
  deploymentEngine: string;
}

export const Diagnostics = () => {

  const [diagnostics, setDiagnostics] = React.useState<DiagnosticsRespoonse>({ deploymentEngine: "" });

  useEffect(() => {
    (async () => {
      const backendUrl = AppConstants.baseUrl;
      const response = await fetch(`${backendUrl}/api/diagnostics`, {
        headers: {
          Accept: 'application/json',
        },
      });

      const result = await response.json();
      setDiagnostics(result);
    })();
  });

  return (<>
    <div className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 className="h2">Diagnostics</h1>
      <div className="btn-toolbar mb-2 mb-md-0">
        <div className="btn-group me-2">
          <button type="button" className="btn btn-sm btn-outline-secondary">Share</button>
          <button type="button" className="btn btn-sm btn-outline-secondary">Export</button>
        </div>
      </div>
    </div>
    <div style={{ whiteSpace: "pre-wrap", fontSize: 11, paddingBottom: 5 }}>
      {diagnostics.deploymentEngine}
    </div>
  </>)
}

export default Diagnostics;