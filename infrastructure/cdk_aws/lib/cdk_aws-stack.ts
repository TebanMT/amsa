import * as cdk from '@aws-cdk/core';
import * as lambda from '@aws-cdk/aws-lambda';
import * as apigateway from '@aws-cdk/aws-apigateway';

export class AmsasStack extends cdk.Stack {

  public readonly authFunction: lambda.Function;
  public readonly createClientFunction: lambda.Function;
  public readonly getAllClientsFunction: lambda.Function;
  public readonly deactivateClientFunction: lambda.Function;
  public readonly createUserFunction: lambda.Function;
  public readonly updateClientFunction: lambda.Function;
  public readonly createSuppliers: lambda.Function;
  public readonly getSuppliers: lambda.Function;
  public readonly updateSuppliers: lambda.Function;
  public readonly deactivateSuppliers: lambda.Function;
  public readonly createProduct: lambda.Function;
  public readonly getProducts: lambda.Function;
  public readonly updateProduct: lambda.Function;
  public readonly deactivateProduct: lambda.Function;
  public readonly getCategories: lambda.Function;
  public readonly getUnits: lambda.Function;
  public readonly getCountSuppliers: lambda.Function;
  public readonly getCountProducts: lambda.Function;
  public readonly getCountClients: lambda.Function;
  public readonly createSale: lambda.Function;
  public readonly getSales: lambda.Function;
  public readonly createBill: lambda.Function;
  public readonly existBillByFolio: lambda.Function;
  public readonly vincularFactura: lambda.Function;
  public readonly existPurchaseByFolio: lambda.Function;
  public readonly createPurchaseOrder: lambda.Function;
  public readonly getAllPurchaseOrders: lambda.Function;
  public readonly getAllBills: lambda.Function;
  public readonly vincularPurchaseOrder: lambda.Function;

  

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);


    

    this.authFunction = new lambda.Function(this, 'authLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/auth/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.createClientFunction = new lambda.Function(this, 'createClientLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/client/create/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getAllClientsFunction = new lambda.Function(this, 'GetAllClientsLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/client/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.deactivateClientFunction = new lambda.Function(this, 'DeactivateClientLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/client/deactivate/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getCountClients = new lambda.Function(this, 'GetCountClientsLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/client/get_count/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });




    this.createUserFunction = new lambda.Function(this, 'CreateUserLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/user/create/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.updateClientFunction = new lambda.Function(this, 'UpdateClientLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/client/update/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.createSuppliers = new lambda.Function(this, 'CreateSupplierLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/supplier/create/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getSuppliers = new lambda.Function(this, 'GetSupplierLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/supplier/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.updateSuppliers = new lambda.Function(this, 'UpdateSupplierLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/supplier/update/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.deactivateSuppliers = new lambda.Function(this, 'DeactivateSupplierLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/supplier/deactivate/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getCountSuppliers = new lambda.Function(this, 'GetCountSuppliersLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/supplier/get_count/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });





    this.createProduct = new lambda.Function(this, 'CreateProductLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/product/create/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getProducts = new lambda.Function(this, 'GetProductsLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/product/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.updateProduct = new lambda.Function(this, 'UpdateProductLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/product/update/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.deactivateProduct = new lambda.Function(this, 'DeactivateProductsLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/product/deactivate/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getCountProducts = new lambda.Function(this, 'GetCountProductsLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/product/get_count/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });






    this.getCategories = new lambda.Function(this, 'GetCategoriesLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/category/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getUnits = new lambda.Function(this, 'GetUnitsLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/unit/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });





    //SALES
    this.createSale = new lambda.Function(this, 'CreateSaleLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/sales/create/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getSales = new lambda.Function(this, 'GetSalesLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/sales/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });


    //BILLS
    this.createBill = new lambda.Function(this, 'CreateBillLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/bills/create/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.existBillByFolio = new lambda.Function(this, 'ExistBillByFolioLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/bills/exist_by_folio/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.vincularFactura = new lambda.Function(this, 'VincluarFactura', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/bills/vincular/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getAllBills = new lambda.Function(this, 'GetAllBillsLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/bills/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });


    //PURCHASE ORDERS
    this.createPurchaseOrder = new lambda.Function(this, 'CreatePurchaseOrderLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/purchase_orders/create/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.existPurchaseByFolio = new lambda.Function(this, 'ExistPurchaseByFolioLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/purchase_orders/exist_by_folio/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.getAllPurchaseOrders = new lambda.Function(this, 'GetAllPurchaseOrdersLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/purchase_orders/get_all/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });

    this.vincularPurchaseOrder = new lambda.Function(this, 'VincluarPurchaseOrderLambda', {
      runtime: lambda.Runtime.PROVIDED_AL2, // Especifica el entorno de ejecución
      code: lambda.Code.fromAsset('../../interfaces/lambda_handlers/purchase_orders/vincular/bootstrap.zip'), // Directorio de tu código Lambda
      handler: 'bootstrap', // Nombre del archivo y método handler
    });





    const api = new apigateway.RestApi(this, 'AmsaApiRest', {
      restApiName: 'AMSA API REST',
      description: 'This API serves enpoints to be consume by AMSA app.',
      defaultCorsPreflightOptions: {
        allowOrigins: apigateway.Cors.ALL_ORIGINS,
        allowMethods: apigateway.Cors.ALL_METHODS,
        allowHeaders: ['Content-Type', 'X-Amz-Date', 'Authorization', 'X-Api-Key'],
      }
    });
    
    const lambdaIntegrationAuth = new apigateway.LambdaIntegration(this.authFunction);
    const resourceAuth = api.root.addResource('auth');
    resourceAuth.addMethod('GET', lambdaIntegrationAuth);

    const lambdaIntegrationCreateClient = new apigateway.LambdaIntegration(this.createClientFunction);
    const resourceCreateClient = api.root.addResource('client');
    resourceCreateClient.addMethod('POST', lambdaIntegrationCreateClient); 

    const lambdaIntegrationGetAllClients = new apigateway.LambdaIntegration(this.getAllClientsFunction);
    resourceCreateClient.addMethod('GET', lambdaIntegrationGetAllClients); 

    const lambdaDeactivateClient = new apigateway.LambdaIntegration(this.deactivateClientFunction);
    const resourceDeactivateClient= resourceCreateClient.addResource('{idClient}');
    resourceDeactivateClient.addMethod('DELETE', lambdaDeactivateClient);

    const lambdaUpdateClient = new apigateway.LambdaIntegration(this.updateClientFunction);
    resourceDeactivateClient.addMethod('PUT', lambdaUpdateClient);

    const lambdaIntegrationCountClients = new apigateway.LambdaIntegration(this.getCountClients);
    const resourceCountClient = resourceCreateClient.addResource('count');
    resourceCountClient.addMethod('GET', lambdaIntegrationCountClients);







    const lambdaIntegrationCreateUser = new apigateway.LambdaIntegration(this.createUserFunction);
    const resourceCreateUser = api.root.addResource('user');
    resourceCreateUser.addMethod('POST', lambdaIntegrationCreateUser);


    const lambdaIntegrationCreateSupplier = new apigateway.LambdaIntegration(this.createSuppliers);
    const resourceSuppliers = api.root.addResource('suppliers');
    resourceSuppliers.addMethod('POST', lambdaIntegrationCreateSupplier); 

    const lambdaIntegrationGetSuppliers = new apigateway.LambdaIntegration(this.getSuppliers);
    const resourceGetSuppliers = resourceSuppliers.addResource('{idSupplier}');
    resourceSuppliers.addMethod('GET', lambdaIntegrationGetSuppliers);

    const lambdaIntegrationUpdateSuppliers = new apigateway.LambdaIntegration(this.updateSuppliers);
    resourceGetSuppliers.addMethod('PUT', lambdaIntegrationUpdateSuppliers); 

    const lambdaIntegrationDeactivateSuppliers = new apigateway.LambdaIntegration(this.deactivateSuppliers);
    resourceGetSuppliers.addMethod('DELETE', lambdaIntegrationDeactivateSuppliers);

    const lambdaIntegrationCountSupliers = new apigateway.LambdaIntegration(this.getCountSuppliers);
    const resourceCountSupplier = resourceSuppliers.addResource('count');
    resourceCountSupplier.addMethod('GET', lambdaIntegrationCountSupliers);





    const lambdaIntegrationCreateProduct = new apigateway.LambdaIntegration(this.createProduct);
    const resourceProducts = api.root.addResource('products');
    resourceProducts.addMethod('POST', lambdaIntegrationCreateProduct); 

    const lambdaIntegrationGetProducts = new apigateway.LambdaIntegration(this.getProducts);
    const resourceGetProducts = resourceProducts.addResource('{idProduct}');
    resourceProducts.addMethod('GET', lambdaIntegrationGetProducts);

    const lambdaIntegrationUpdateProduct = new apigateway.LambdaIntegration(this.updateProduct);
    resourceGetProducts.addMethod('PUT', lambdaIntegrationUpdateProduct); 

    const lambdaIntegrationDeactivateProduct = new apigateway.LambdaIntegration(this.deactivateProduct);
    resourceGetProducts.addMethod('DELETE', lambdaIntegrationDeactivateProduct);

    const lambdaIntegrationCountProducts = new apigateway.LambdaIntegration(this.getCountProducts);
    const resourceCountProduct = resourceProducts.addResource('count');
    resourceCountProduct.addMethod('GET', lambdaIntegrationCountProducts);



  






    const lambdaIntegrationGetCategories = new apigateway.LambdaIntegration(this.getCategories);
    const resourceCategory = api.root.addResource('categories');
    resourceCategory.addMethod('GET', lambdaIntegrationGetCategories);

    const lambdaIntegrationGetUnits = new apigateway.LambdaIntegration(this.getUnits);
    const resourceUnit = api.root.addResource('units');
    resourceUnit.addMethod('GET', lambdaIntegrationGetUnits); 


    //SALES
    const lambdaIntegrationCreateSale = new apigateway.LambdaIntegration(this.createSale);
    const resourceSale = api.root.addResource('sales');
    resourceSale.addMethod('POST', lambdaIntegrationCreateSale);
  
    const lambdaIntegrationGetSales = new apigateway.LambdaIntegration(this.getSales);
    resourceSale.addMethod('GET', lambdaIntegrationGetSales);



    //BILLS
    const lambdaIntegrationCreateBill = new apigateway.LambdaIntegration(this.createBill);
    const resourceBill = api.root.addResource('bills');
    resourceBill.addMethod('POST', lambdaIntegrationCreateBill);

    const lambdaIntegrationGetAllBills = new apigateway.LambdaIntegration(this.getAllBills);
    resourceBill.addMethod('GET', lambdaIntegrationGetAllBills);

    const lambdaIntegrationExistBillByFolio = new apigateway.LambdaIntegration(this.existBillByFolio);
    const resourceExistFolio = resourceBill.addResource('exist');
    resourceExistFolio.addMethod('GET', lambdaIntegrationExistBillByFolio);

    const lambdaIntegrationVincular = new apigateway.LambdaIntegration(this.vincularFactura);
    resourceBill.addMethod('PUT', lambdaIntegrationVincular);


    //PURCHASE ORDERS
    const lambdaIntegrationCreatePurchase = new apigateway.LambdaIntegration(this.createPurchaseOrder);
    const resourcePurchaseOrder = api.root.addResource('purchaseOrder');
    resourcePurchaseOrder.addMethod('POST', lambdaIntegrationCreatePurchase);

    const lambdaIntegrationExistPurchaseOrderByFolio = new apigateway.LambdaIntegration(this.existPurchaseByFolio);
    const resourcePurchaseOrderExist = resourcePurchaseOrder.addResource('exist');
    resourcePurchaseOrderExist.addMethod('GET', lambdaIntegrationExistPurchaseOrderByFolio);

    const lambdaIntegrationGetAllPurchaseOrder = new apigateway.LambdaIntegration(this.getAllPurchaseOrders);
    resourcePurchaseOrder.addMethod('GET', lambdaIntegrationGetAllPurchaseOrder);

    const lambdaIntegrationVincularPurchase = new apigateway.LambdaIntegration(this.vincularPurchaseOrder);
    resourcePurchaseOrder.addMethod('PUT', lambdaIntegrationVincularPurchase);


  }




}
