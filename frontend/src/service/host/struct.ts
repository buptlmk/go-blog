export class HostEntity {
    id: number = 0;
    name: string = "";
    ip: string = "";
    memory: number = 0;
    cpu: number = 0;
    gpu: Array<number> = [];
    disk: number = 0;
    date: string = "";
    nfs_location: string = ""
}


export interface GpuProcessUsageEntity {
    pid: number;
    type: string;
    process_name: string;
    used_memory: number;
}

export interface GpuUsageEntity {
    name: string
    minor_number: number
    uuid: string //"GPU-11d5fae1-6870-f322-48e7-b82c2105f1e5";
    gpu_util: number //0;
    memory_util: number //0;
    fan_speed: number //0;
    temperature: number //39;
    power_state: string //"P8";
    power_draw: number //21.75;
    power_limit: number //260;
    memory_total: number //11554258944;
    memory_used: number //0;
    memory_free: number //11554258944;
    processes: Map<string, GpuProcessUsageEntity>
}

export class GpusUsageEntity {
    driver_version: string = "430.26";
    cuda_version: string = "10.2";
    time: string = "2019-08-29T08:10:27.180940909Z";
    gpus: Map<string, GpuUsageEntity> = new Map();
}

export class CoreUsageEntity {
    id: number = 0;
    user: number = 0;
    nice: number = 0;
    system: number = 0;
    idle: number = 0;
    iowait: number = 0;
    irq: number = 0;
    softirq: number = 0;
    steal: number = 0;
    guest: number = 0;
    guest_nice: number = 0;
    total: number = 0;
}

export class MemoryUsageEntity {
    total: number = 0;
    used: number = 0;
    buffers: number = 0;
    cached: number = 0;
    swap_total: number = 0;
    swap_used: number = 0;
}

export class NFSStatusEntity {
    host_id: number = 0;
    online: string = "";
    error: string = "";
    options: string = "";
    source: string = "";
    destine: string = ""
}

export class ResourcesEntity {
    name: string = "";
    ip: string = "";
    memory: number = 0;
    cpu: number = 0;
    disk: number = 0;
    gpu_weight: number = 0;
    gpu: Array<number> = []
}


export class Flow {
    bytes: number = 0;
    packets: number = 0;
    errs: number = 0;
    drop: number = 0;
    fifo: number = 0;
    frame: number = 0;
    compressed: number = 0;
    multicast: number = 0;
}


export class NetInterface {
    receive: Flow = new Flow();
    transmit: Flow = new Flow();
}

export class DiskIO {
    major_number: number = 0;
    minor_number: number = 0;
    reads_completed_times: number = 0;
    reads_merged_times: number = 0;
    sectors_read_times: number = 0;
    time_spent_reading: number = 0;
    writes_completed_times: number = 0;
    writes_merged_times: number = 0;
    sectors_written_times: number = 0;
    time_spent_writing: number = 0;
    ios_currently_in_progress: number = 0;
    time_spent_on_ios: number = 0;
    weighted_time_spent_on_ios: number = 0;
    read_speed: number = 0;
    write_speed: number = 0
}
