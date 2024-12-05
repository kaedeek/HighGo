import cv2, sys
from cv2 import dnn_superres

def super_resolve(input_image, model_path, output_image):
    sr = dnn_superres.DnnSuperResImpl_create()
    sr.readModel(model_path)
    sr.setModel("espcn", 4)
    image = cv2.imread(input_image)
    result = sr.upsample(image)
    cv2.imwrite(output_image, result)

if __name__ == "__main__":
    super_resolve(sys.argv[1], sys.argv[2], sys.argv[3])