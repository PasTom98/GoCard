export interface SystemMetrics {
    goVersion: string;
    numCPU: number;
    numGoroutine: number;
    memoryStats: {
        alloc: number;
        totalAlloc: number;
        sys: number;
        numGC: number;
    };
}

export interface DatabaseMetrics {
    maxOpenConnections: number;
    openConnections: number;
    inUseConnections: number;
    idleConnections: number;
    waitCount: number;
    waitDuration: number;
}

export interface HealthResponse {
    status: string;
    timestamp: string;
    system: SystemMetrics;
    database?: DatabaseMetrics;
}

export async function getHealth(): Promise<HealthResponse> {
    const res = await fetch('/api/0.0.1/health');
    if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
    return res.json();
}
