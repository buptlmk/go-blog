import {Response} from "./response";
import {BaseService} from "./base-service";

export class ImageEntity {
    image_id: string = "";
    image_tag: string = "";
    size: number = 0;
    is_public: boolean = false;
    date: string = ""
}

export class DockerImageService extends BaseService {
    private static instance: DockerImageService;

    public static get = () => {
        if (DockerImageService.instance) {
            return DockerImageService.instance
        } else {
            DockerImageService.instance = new DockerImageService();
            return DockerImageService.instance;
        }
    };

    getImageList(isByUser: boolean = false): Promise<Response<Array<ImageEntity>>> {
        let url = isByUser ? "/api/docker/image/list" : "/api/docker/image/list/public";
        return this.cachedPost<Array<ImageEntity>>(url, {}, 0)
    }

    deleteImage(imageName: string): Promise<Response<Array<ImageEntity>>> {
        return this.cachedPost<Array<ImageEntity>>("/api/docker/image/delete", {image_name: imageName}, 0)
    }

    changeImagePublicState(imageName: string, isPublic: boolean): Promise<Response<boolean>> {
        return this.cachedPost<boolean>("/api/docker/image/public", {image_name: imageName, is_public: isPublic}, 0)
    }

    copyImage(imagePureTag: string, schoolCard: string): Promise<Response<boolean>> {
        return this.cachedPost<boolean>("/api/docker/image/copy", {
            image_pure_tag: imagePureTag,
            school_card: schoolCard
        }, 0)
    }

    autoRoundSize(a: number, d: number = 2) {
        if (0 === a) return "0 Bytes";
        let c = 1024, e = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"],
            f = Math.floor(Math.log(a) / Math.log(c));
        return parseFloat((a / Math.pow(c, f)).toFixed(d)) + " " + e[f]
    }

    parseImageTag(imageTag: string) {
        let hubIndex = imageTag.indexOf('/');
        let tagIndex = imageTag.lastIndexOf(':');
        let hubName = imageTag.substring(0, hubIndex);
        let imageName = imageTag.substring(hubIndex + 1, tagIndex);
        let tag = imageTag.substring(tagIndex + 1, imageTag.length);
        return [hubName, imageName, tag]
    }
}



